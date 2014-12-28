package redis

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/garyburd/redigo/redis"
)

type Client struct {
	P    *redis.Pool
	conn string
}

var redisClient *Client

func NewClient() *Client {
	return &Client{}
}

func (r *Client) Initialize(config string) error {
	var cf map[string]string
	json.Unmarshal([]byte(config), &cf)

	if _, ok := cf["conn"]; !ok {
		return errors.New("config has no conn key")
	}

	r.conn = cf["conn"]
	r.connect()

	c := r.P.Get()
	defer c.Close()

	return c.Err()
}

func (r *Client) connect() {
	dialFunc := func() (c redis.Conn, err error) {
		c, err = redis.Dial("tcp", r.conn)
		return
	}
	r.P = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 180 * time.Second,
		Dial:        dialFunc,
	}
}

func (r *Client) Do(commandName string, args ...interface{}) (reply interface{}, err error) {
	c := r.P.Get()
	defer c.Close()

	return c.Do(commandName, args...)
}

func (r *Client) Get(key string) interface{} {
	if v, err := r.Do("GET", key); err == nil {
		return v
	}
	return nil
}

func (r *Client) Rpush(key string, val interface{}) (err error) {
	_, err = r.Do("RPUSH", key, val)
	return
}

func init() {
	redisClient = NewClient()
}
