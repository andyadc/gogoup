package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*

 */

// https://github.com/garyburd/redigo

var (
	Pool *redis.Pool
)

func init() {
	redisHost := `www.jd-server.com:6379`
	Pool = newPool(redisHost)
	close()
}

func main() {
	test, err := get("a")
	fmt.Println(test, err)
	fmt.Println(string(test), err)
}

func get(key string) ([]byte, error) {
	conn := Pool.Get()
	defer conn.Close()

	data, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return data, fmt.Errorf("error get key %s: %v", key, err)
	}
	return data, err
}

func newPool(server string) *redis.Pool {

	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", server)
			if err != nil {
				return nil, err
			}
			if _, err := c.Do("AUTH", "andy"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
}

func close() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, syscall.SIGTERM)
	signal.Notify(c, syscall.SIGKILL)
	go func() {
		<-c
		Pool.Close()
		os.Exit(0)
	}()
}
