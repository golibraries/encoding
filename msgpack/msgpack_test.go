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
		data, err := msgpack.Marshal(&Access{
			Type:    1,
			Account: []byte("123456"),
			Enable:  true,
		})
		c.Assert(err, IsNil)

		access := &Access{}
		err = msgpack.Unmarshal(data, access)
		c.Assert(err, IsNil)
		c.Assert(access.Type, Equals, int64(1))
		c.Assert(access.Enable, Equals, true)
		c.Assert(access.Account, DeepEquals, []byte("123456"))
	})
}
