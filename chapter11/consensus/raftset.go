package consensus

import (
	"fmt"

	"github.com/hashicorp/raft"
)

// keep a map of rafts for later
var rafts map[raft.ServerAddress]*raft.Raft

func init() {
	rafts = make(map[raft.ServerAddress]*raft.Raft)
}

// raftSet stores all the setup material we need
type raftSet struct {
	Config        *raft.Config
	Store         *raft.InmemStore
	SnapShotStore raft.SnapshotStore
	FSM           *FSM
	Transport     raft.LoopbackTransport
	Configuration raft.Configuration
}

// generate n raft sets to bootstrap the raft cluster
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

	// configuration needs to be consistent between
	// services and so we need the full serverlist in this
	// case
	for _, r := range rs {
		r.Configuration = raft.Configuration{Servers: servers}
	}

	return rs
}
