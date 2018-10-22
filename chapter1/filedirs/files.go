package filedirs

import (
	"bytes"
	"io"
	"os"
	"strings"
)

// Capitalizer opens a file, reads the contents,
// then writes those contents to a second file
func Capitalizer(f1 *os.File, f2 *os.File) error {
	if _, err := f1.Seek(0, 0); err != nil {
		return err
	}

	var tmp = new(bytes.Buffer)

	if _, err := io.Copy(tmp, f1); err != nil {
		return err
	}

	s := strings.ToUpper(tmp.String())

	if _, err := io.Copy(f2, strings.NewReader(s)); err != nil {
		return err
	}
	return nil
}

// CapitalizerExample creates two files, writes to one
//then calls Capitalizer() on both
func CapitalizerExample() error {
	f1, err := os.Create("file1.txt")
	if err != nil {
		return err
	}

	if _, err := f1.Write([]byte(`
    this file contains
    a number of words
    and new lines`)); err != nil {
		return err
	}

	f2, err := os.Create("file2.txt")
	if err != nil {
		return err
	}

	if err := Capitalizer(f1, f2); err != nil {
		return err
	}

	if err := os.Remove("file1.txt"); err != nil {
		return err
	}

	if err := os.Remove("file2.txt"); err != nil {
		return err
	}

	return nil
}
