package synchronizer

import (
	"context"
	"sync"
	"time"

	"github.com/relab/hotstuff"
	"github.com/relab/hotstuff/internal/logging"
)

var logger = logging.GetLogger()

type Synchronizer struct {
	hotstuff.LeaderRotation

	mut      sync.Mutex
	lastBeat hotstuff.View
	timeout  time.Duration
	timer    *time.Timer
	stop     context.CancelFunc
	hs       hotstuff.Consensus
	stopped  bool
}

func New(leaderRotation hotstuff.LeaderRotation, initialTimeout time.Duration) *Synchronizer {
	return &Synchronizer{
		LeaderRotation: leaderRotation,
		timeout:        initialTimeout,
	}
}

// OnPropose should be called when a replica has received a new valid proposal.
func (s *Synchronizer) OnPropose() {
	s.mut.Lock()
	defer s.mut.Unlock()
	if s.timer != nil {
		s.timer.Reset(s.timeout)
	}
}

// OnFinishQC should be called when a replica has created a new qc
func (s *Synchronizer) OnFinishQC() {
	s.beat()
}

// OnNewView should be called when a replica receives a valid NewView message
func (s *Synchronizer) OnNewView() {
	s.beat()
}

func (s *Synchronizer) Init(hs hotstuff.Consensus) {
	s.hs = hs
}

// Start starts the synchronizer
func (s *Synchronizer) Start() {
	if s.GetLeader(s.hs.Leaf().View()+1) == s.hs.Config().ID() {
		s.hs.Propose()
	}
	s.timer = time.NewTimer(s.timeout)
	var ctx context.Context
	ctx, s.stop = context.WithCancel(context.Background())
	go s.newViewTimeout(ctx)
}

// Stop stops the synchronizer
func (s *Synchronizer) Stop() {
	s.stopped = true
	if s.timer != nil && !s.timer.Stop() {
		<-s.timer.C
	}
	s.stop()
}

func (s *Synchronizer) beat() {
	if s.stopped {
		return
	}
	view := s.hs.Leaf().View()
	s.mut.Lock()
	if view <= s.lastBeat {
		s.mut.Unlock()
		logger.Debug("Can't beat more than once per view ", s.lastBeat)
		return
	}
	if s.GetLeader(view+1) != s.hs.Config().ID() {
		s.mut.Unlock()
		return
	}
	s.lastBeat = view
	s.mut.Unlock()
	s.hs.Propose()
}

func (s *Synchronizer) newViewTimeout(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-s.timer.C:
			s.hs.CreateDummy()
			s.hs.NewView()
			s.mut.Lock()
			s.timer.Reset(s.timeout)
			s.mut.Unlock()
		}
	}
}
