# Gremlin

After watching all of Liz Rice's [Containers from Scratch](https://github.com/lizrice/containers-from-scratch) videos, I
realised that these containers concepts are relatively old sysadmin knowledge that most backend engineers should already
use and understand. Yet, not many of my coworkers actually understand what "containers are just a chroot jailed process"
means.

This is a relatively harmless Golang process that will constantly attempt to carry out host harassment. Only after
implementing your own runtime with Liz's knowledge can you contain and tame the process.

## Prerequisites

* Golang
* Linux VM. Or YOLO on your Linux host machine.
* No existing directories called `/food` or `/poop`.
* Okay with running random code of mine with `sudo`.

## Instructions

1. `git clone git@github.com:alienchow/gremlin.git`
2. `cd gremlin`
3. `go build -o gremlin cmd/gremlin/*.go`
4. Run the `gremlin` executable binary with whatever your solution is.

