package consensus

import (
	"io"

	"github.com/hashicorp/raft"
)

// FSM implements the raft FSM interface
// and holds a state
type FSM struct {
	state state
}

// NewFSM creates a new FSM with
// start state of "first"
func NewFSM() *FSM {
	return &FSM{state: first}
}

// Apply updates our FSM
func (f *FSM) Apply(r *raft.Log) interface{} {
	f.state.Transition(state(r.Data))
	return string(f.state)
}

// Snapshot needed to satisfy the raft FSM interface
func (f *FSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

// Restore needed to satisfy the raft FSM interface
func (f *FSM) Restore(io.ReadCloser) error {
	return nil
}
