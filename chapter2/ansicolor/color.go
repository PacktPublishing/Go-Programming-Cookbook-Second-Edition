package ansicolor

import "fmt"

//Color of text
type Color int

const (
	// ColorNone is default
	ColorNone = iota
	// Red colored text
	Red
	// Green colored text
	Green
	// Yellow colored text
	Yellow
	// Blue colored text
	Blue
	// Magenta colored text
	Magenta
	// Cyan colored text
	Cyan
	// White colored text
	White
	// Black colored text
	Black Color = -1
)

// ColorText holds a string and its color
type ColorText struct {
	TextColor Color
	Text      string
}

func (r *ColorText) String() string {
	if r.TextColor == ColorNone {
		return r.Text
	}

	value := 30
	if r.TextColor != Black {
		value += int(r.TextColor)
	}
	return fmt.Sprintf("\033[0;%dm%s\033[0m", value, r.Text)
}
