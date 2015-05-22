package accessctl

import (
	"errors"
	"reflect"

	"github.com/ikeikeikeike/gopkg/convert"
	"github.com/ikeikeikeike/gopkg/redis"
)

type Counter struct {
	rc      *redis.Client
	baseKey string
}

func NewCounter(config string) (*Counter, error) {
	rc := redis.NewClient()

	err := rc.Initialize(config)
	if err != nil {
		return nil, err
	}

	return &Counter{
		rc:      rc,
		baseKey: "accessctl.counter",
	}, nil
}

func (c *Counter) PushUA(ua string) error {
	if !IsUA(ua) {
		return errors.New("The user agent is a not allowed.")
	}
	return c.Push(ua)
}

func (c *Counter) Push(value string) error {
	return c.rc.Rpush(c.baseKey, value)
}

func (c *Counter) Listall() (list []string) {
	s := reflect.ValueOf(c.rc.Listall(c.baseKey))

	list = make([]string, s.Len())
	for i := 0; i < s.Len(); i++ {
		list[i] = convert.ToStr(s.Index(i).Elem().Bytes())
	}
	return
}

func (c *Counter) Clean() {
	c.rc.Del(c.baseKey)
}
