#!/bin/bash

image="hotstuff_worker"

# ensure that the image is built
docker images | grep "$image" &>/dev/null \
	|| docker build -t "$image" -f "./Dockerfile.worker" ".."

container="$(docker run --rm -d -p 2020:22 -p 4000:4000 "$image")"

sleep 1s

ssh-keyscan -p 2020 localhost > known_hosts

docker logs --follow "$container"
docker rm -f "$container" &>/dev/null
