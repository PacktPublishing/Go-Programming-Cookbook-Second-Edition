package mocking

type MockDoStuffer struct {
	// closure to assist with mocking
	MockDoStuff func(input string) error
}

func (m *MockDoStuffer) DoStuff(input string) error {
	if m.MockDoStuff != nil {
		return m.MockDoStuff(input)
	}
	// if we don't mock, return a common case
	return nil
}
