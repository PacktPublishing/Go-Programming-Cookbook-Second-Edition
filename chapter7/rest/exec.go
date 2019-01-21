package rest

import "fmt"

// Exec creates an API Client and uses its
// GetGoogle method, then prints the result
func Exec() error {
	c := NewAPIClient("username", "password")

	StatusCode, err := c.GetGoogle()
	if err != nil {
		return err
	}
	fmt.Println("Result of GetGoogle:", StatusCode)
	return nil
}
