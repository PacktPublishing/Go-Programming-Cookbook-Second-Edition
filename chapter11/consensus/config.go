package consensus

import (
	"github.com/hashicorp/raft"
)

// Config creates num in-memory raft
// nodes and connects them
func Config(num int) {

	// create n "raft-sets" consisting of
	// everything needed to represent a node
	rs := getRaftSet(num)

	//connect all of the transports
	for _, r1 := range rs {
		for _, r2 := range rs {
			r1.Transport.Connect(r2.Transport.LocalAddr(), r2.Transport)
		}
	}

	// for each node, bootstrap then connect
	for _, r := range rs {
		if err := raft.BootstrapCluster(r.Config, r.Store, r.Store, r.SnapShotStore, r.Transport, r.Configuration); err != nil {
			panic(err)
		}
		raft, err := raft.NewRaft(r.Config, r.FSM, r.Store, r.Store, r.SnapShotStore, r.Transport)
		if err != nil {
			panic(err)
		}
		rafts[r.Transport.LocalAddr()] = raft
	}
}
