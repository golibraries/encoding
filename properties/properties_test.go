package properties_test

import (
	"testing"

	"github.com/golibraries/encoding/properties"

	. "github.com/frankban/quicktest"
)

type Access struct {
	Type    int64  `properties:"type"`
	Account string `properties:"account"`
	Enable  bool   `properties:"enable"`
}

func TestProperties(t *testing.T) {
	c := New(t)
	c.Run("TestProperties", func(c *C) {
		data, err := properties.Marshal(&Access{
			Type:    1,
			Account: "123456",
			Enable:  true,
		})
		c.Assert(err, IsNil)
		c.Assert(string(data), Equals, `account = 123456
enable = true
type = 1
`)

		access := &Access{}
		err = properties.Unmarshal(data, access)
		c.Assert(err, IsNil)
		c.Assert(access.Type, Equals, int64(1))
		c.Assert(access.Account, DeepEquals, "123456")
		c.Assert(access.Enable, Equals, true)
	})
}
