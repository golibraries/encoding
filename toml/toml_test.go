package toml_test

import (
	"testing"

	"github.com/golibraries/encoding/toml"

	. "github.com/frankban/quicktest"
)

type Access struct {
	Type    int64  `toml:"type"`
	Account string `toml:"account"`
	Enable  bool   `toml:"enable"`
}

func TestTOML(t *testing.T) {
	c := New(t)
	c.Run("TestTOML", func(c *C) {
		data, err := toml.Marshal(&Access{
			Type:    1,
			Account: "123456",
			Enable:  true,
		})
		c.Assert(err, IsNil)
		c.Assert(string(data), Equals, `type = 1
account = '123456'
enable = true
`)

		access := &Access{}
		err = toml.Unmarshal(data, access)
		c.Assert(err, IsNil)
		c.Assert(access.Type, Equals, int64(1))
		c.Assert(access.Account, DeepEquals, "123456")
		c.Assert(access.Enable, Equals, true)
	})
}
