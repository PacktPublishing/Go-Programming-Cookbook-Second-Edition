package discovery

import "fmt"

// Exec creates a consul entry then queries it
func Exec(cli Client) error {
	if err := cli.Register([]string{"Go", "Awesome"}); err != nil {
		return err
	}

	entries, _, err := cli.Service("discovery", "Go")
	if err != nil {
		return err
	}
	for _, entry := range entries {
		fmt.Printf("%#v\n", entry.Service)
	}

	return nil
}
