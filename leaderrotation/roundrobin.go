package leaderrotation

import (
	"sort"

	"github.com/relab/hotstuff"
	"github.com/relab/hotstuff/modules"
)

func init() {
	modules.RegisterModule("round-robin", NewRoundRobin)
}

type roundRobin struct {
	configuration modules.Configuration
	replicas      []hotstuff.ID
}

func (rr *roundRobin) InitModule(mods *modules.Core) {
	mods.Get(&rr.configuration)
}

// GetLeader returns the id of the leader in the given view
func (rr roundRobin) GetLeader(view hotstuff.View) hotstuff.ID {
	// TODO: does not support reconfiguration
	// assume IDs start at 1
	if rr.replicas == nil {
		rr.replicas = make([]hotstuff.ID, 0, rr.configuration.Len())
		for id := range rr.configuration.Replicas() {
			rr.replicas = append(rr.replicas, id)
		}
		sort.Slice(rr.replicas, func(i, j int) bool {
			return rr.replicas[i] < rr.replicas[j]
		})
	}
	idx := view % hotstuff.View(len(rr.replicas))
	id := rr.replicas[idx]
	// rr.logger.Infof("view %d: leader is %d", view, id)
	return hotstuff.ID(id)
}

// NewRoundRobin returns a new round-robin leader rotation implementation.
func NewRoundRobin() modules.LeaderRotation {
	return &roundRobin{}
}

func chooseRoundRobin(view hotstuff.View, numReplicas int) hotstuff.ID {
	return hotstuff.ID(view%hotstuff.View(numReplicas) + 1)
}
