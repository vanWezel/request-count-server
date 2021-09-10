package redis

import (
	"github.com/gomodule/redigo/redis"
)

type Client struct {
	Interface
	pool *redis.Pool
	db   int
}

func New(host string, db int) *Client {
	return &Client{
		db: db,
		pool: &redis.Pool{
			MaxIdle:     0,
			MaxActive:   0,
			IdleTimeout: 0,
			Wait:        false,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", host)
			},
		},
	}
}

func (c *Client) getConn() (redis.Conn, error) {
	conn := c.pool.Get()
	if _, err := conn.Do("SELECT", c.db); err != nil {
		_ = conn.Close()
		return nil, err
	}

	return conn, nil
}

func (c *Client) Ping() error {
	conn, err := c.getConn()
	if err != nil {
		return err
	}

	return conn.Send("PING")
}

func (c *Client) FlushDb() error {
	conn, err := c.getConn()
	if err != nil {
		return err
	}

	_, err = conn.Do("FLUSHDB")

	return err
}

func (c *Client) GetInt(key string) (int, error) {
	conn, err := c.getConn()
	if err != nil {
		return 0, err
	}

	return redis.Int(conn.Do("GET", key))
}

func (c *Client) Increment(key string) (int, error) {
	conn, err := c.getConn()
	if err != nil {
		return 0, err
	}

	return redis.Int(conn.Do("INCR", key))
}
