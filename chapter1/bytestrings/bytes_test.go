package bytestrings

import (
	"testing"
)

func TestWorkWithBuffer(t *testing.T) {
	err := WorkWithBuffer()
	if err != nil {
		t.Errorf("unexpected error")
	}
}
