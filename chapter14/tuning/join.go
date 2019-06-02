package tuning

import "strings"

func join(vals ...string) string {
	c := strings.Join(vals, " ")
	return c
}
