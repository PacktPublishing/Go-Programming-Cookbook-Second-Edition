package consensus

import (
	"fmt"

	"github.com/hashicorp/raft"
)

var rafts map[raft.ServerAddress]*raft.Raft

func init() {
	rafts = make(map[raft.ServerAddress]*raft.Raft)
}

type raftSet struct {
	Config        *raft.Config
	Store         *raft.InmemStore
	SnapShotStore raft.SnapshotStore
	FSM           *FSM
	Transport     raft.LoopbackTransport
	Configuration raft.Configuration
}

func getRaftSet(num int) []*raftSet {
	rs := make([]*raftSet, num)
	servers := make([]raft.Server, num)
	for i := 0; i < num; i++ {
		addr := raft.ServerAddress(fmt.Sprint(i))
		_, transport := raft.NewInmemTransport(addr)
		servers[i] = raft.Server{
			Suffrage: raft.Voter,
			ID:       raft.ServerID(addr),
			Address:  addr,
		}
		config := raft.DefaultConfig()
		config.LocalID = raft.ServerID(addr)

		rs[i] = &raftSet{
			Config:        config,
			Store:         raft.NewInmemStore(),
			SnapShotStore: raft.NewInmemSnapshotStore(),
			FSM:           NewFSM(),
			Transport:     transport,
		}
	}

	for _, r := range rs {
		r.Configuration = raft.Configuration{Servers: servers}
	}

	return rs
}

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
