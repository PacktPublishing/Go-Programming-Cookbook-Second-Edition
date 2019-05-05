package consensus

type state string

const (
	first  state = "first"
	second       = "second"
	third        = "third"
)

var allowedState map[state][]state

func init() {
	// setup valid states
	allowedState = make(map[state][]state)
	allowedState[first] = []state{second, third}
	allowedState[second] = []state{third}
	allowedState[third] = []state{first}
}

// CanTransition checks if a new state is valid
func (s *state) CanTransition(next state) bool {
	for _, n := range allowedState[*s] {
		if n == next {
			return true
		}
	}
	return false
}

// Transition will move a state to the next
// state if able
func (s *state) Transition(next state) {
	if s.CanTransition(next) {
		*s = next
	}
}
