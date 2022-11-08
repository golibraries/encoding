package json_test

import (
	"testing"

	"github.com/golibraries/encoding/json"

	. "github.com/frankban/quicktest"
)

type Access struct {
	Type    int64  `json:"type"`
	Account []byte `json:"account"`
	Enable  bool   `json:"enable"`
}

func TestJSON(t *testing.T) {
	c := New(t)
	c.Run("TestJSON", func(c *C) {
		a := &Access{
			Type:    1,
			Account: []byte("123456"),
			Enable:  true,
		}
		b, err := json.Marshal(a)
		c.Assert(err, IsNil)
		c.Assert(string(b), Equals, `{"type":1,"account":"MTIzNDU2","enable":true}`)

		var access Access
		c.Assert(json.Unmarshal(b, &access), IsNil)
		c.Assert(access.Type, Equals, int64(1))
		c.Assert(access.Enable, Equals, true)
		c.Assert(access.Account, DeepEquals, []byte("123456"))
	})
}
