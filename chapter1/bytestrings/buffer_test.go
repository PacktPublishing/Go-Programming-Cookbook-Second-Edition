package bytestrings

import (
	"bytes"
	"testing"
)

func TestBuffer(t *testing.T) {
	var bufferTests = []struct {
		input  string
		output *bytes.Buffer
	}{
		{"abc", bytes.NewBufferString("abc")},
		{"", bytes.NewBufferString("")},
		{"❤️", bytes.NewBufferString("❤️")},
	}

	for _, tt := range bufferTests {
		if got, want := Buffer(tt.input), tt.output; got.String() != want.String() {
			t.Errorf("got: %#v; want: %#v", got, want)
		}
	}
}

func TestToString(t *testing.T) {
	var toStringTests = []struct {
		input   *bytes.Buffer
		output1 string
		output2 error
	}{
		{bytes.NewBufferString("abc"), "abc", nil},
		//{nil, "", nil},
	}

	for _, tt := range toStringTests {
		got1, got2 := toString(tt.input)
		if got1 != tt.output1 || got2 != tt.output2 {
			t.Errorf("got: %#v %#v; want: %#v %#v", got1, got2, tt.output1, tt.output2)
		}
	}
}
