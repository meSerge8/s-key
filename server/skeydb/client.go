package skeydb

import (
	"server/skey"
)

type Client struct {
	id  uint32
	key skey.Key
}

func (c *Client) ConfirmKey(key skey.Key) bool {
	if !skey.Check(key, c.key) {
		return false
	}
	c.key = key
	return true
}

func (c *Client) GetId() uint32 {
	return c.id
}
