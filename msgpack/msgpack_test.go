package msgpack_test

import (
	"testing"

	. "github.com/frankban/quicktest"
	"github.com/golibraries/encoding/msgpack"
)

type Access struct {
	Type    int64  `msgpack:"type"`
	Account []byte `msgpack:"account"`
	Enable  bool   `msgpack:"enable"`
}

func TestMsgpack(t *testing.T) {
	c := New(t)
	c.Run("TestMsgpack", func(c *C) {
		access := &Access{
			Type:    1,
			Account: []byte("123456"),
			Enable:  true,
		}
		data, err := msgpack.Marshal(access)
		c.Assert(err, IsNil)
		var acess2 Access
		c.Assert(msgpack.Unmarshal(data, &acess2), IsNil)
		c.Assert(acess2.Type, Equals, int64(1))
		c.Assert(acess2.Enable, Equals, true)
		c.Assert(acess2.Account, DeepEquals, []byte("123456"))
	})
}
