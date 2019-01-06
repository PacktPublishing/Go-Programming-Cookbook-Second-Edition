package redis

import (
	"fmt"
	"time"

	redis "gopkg.in/redis.v5"
)

// Exec performs some redis operations
func Exec() error {
	conn, err := Setup()
	if err != nil {
		return err
	}

	c1 := "value"
	// value is an interface, we can store whatever
	// the last argument is the redis expiration
	conn.Set("key", c1, 5*time.Second)

	var result string
	if err := conn.Get("key").Scan(&result); err != nil {
		switch err {
		// this means the key
		// was not found
		case redis.Nil:
			return nil
		default:
			return err
		}
	}

	fmt.Println("result =", result)

	return nil
}
