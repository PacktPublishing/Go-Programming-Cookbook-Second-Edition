package redis

import (
	"fmt"

	redis "gopkg.in/redis.v5"
)

// Sort performs a sort redis operations
func Sort() error {
	conn, err := Setup()
	if err != nil {
		return err
	}

	if err := conn.LPush("list", 1).Err(); err != nil {
		return err
	}
	if err := conn.LPush("list", 3).Err(); err != nil {
		return err
	}
	if err := conn.LPush("list", 2).Err(); err != nil {
		return err
	}

	res, err := conn.Sort("list", redis.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)
	conn.Del("list")
	return nil
}
