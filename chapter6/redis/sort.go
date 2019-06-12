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

	listkey := "list"
	if err := conn.LPush(listkey, 1).Err(); err != nil {
		return err
	}
	// this will clean up the list key if any of the subsequent commands error
	defer conn.Del(listkey)

	if err := conn.LPush(listkey, 3).Err(); err != nil {
		return err
	}
	if err := conn.LPush(listkey, 2).Err(); err != nil {
		return err
	}

	res, err := conn.Sort(listkey, redis.Sort{Order: "ASC"}).Result()
	if err != nil {
		return err
	}
	fmt.Println(res)

	return nil
}
