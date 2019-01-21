package decorator

import "fmt"

// Exec creates a client, calls google.com
// then prints the response
func Exec() error {
	c := Setup()

	resp, err := c.Get("https://www.google.com")
	if err != nil {
		return err
	}
	fmt.Println("Response code:", resp.StatusCode)
	return nil
}
