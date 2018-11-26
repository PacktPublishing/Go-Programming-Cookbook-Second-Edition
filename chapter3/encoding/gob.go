package encoding

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

// pos stores the x, y position
// for Object
type pos struct {
	X      int
	Y      int
	Object string
}

// GobExample demonstrates using
// the gob package
func GobExample() error {
	buffer := bytes.Buffer{}

	p := pos{
		X:      10,
		Y:      15,
		Object: "wrench",
	}

	// note that if p was an interface
	// we'd have to call gob.Register first

	e := gob.NewEncoder(&buffer)
	if err := e.Encode(&p); err != nil {
		return err
	}

	// note this is a binary format so it wont print well
	fmt.Println("Gob Encoded valued length: ", len(buffer.Bytes()))

	p2 := pos{}
	d := gob.NewDecoder(&buffer)
	if err := d.Decode(&p2); err != nil {
		return err
	}

	fmt.Println("Gob Decode value: ", p2)

	return nil
}
