package json_test

import (
	"testing"

	"github.com/golibraries/encoding/json"

	. "github.com/frankban/quicktest"
)

type Access struct {
	Type    int64  `json:"type"`
	Account string `json:"account"`
	Enable  bool   `json:"enable"`
}

func TestJSON(t *testing.T) {
	c := New(t)
	c.Run("TestJSON", func(c *C) {
		data, err := json.Marshal(&Access{
			Type:    1,
			Account: "123456",
			Enable:  true,
		})
		c.Assert(err, IsNil)
		c.Assert(string(data), Equals, `{"type":1,"account":"123456","enable":true}`)

		access := &Access{}
		err = json.Unmarshal(data, access)
		c.Assert(err, IsNil)
		c.Assert(access.Type, Equals, int64(1))
		c.Assert(access.Account, DeepEquals, "123456")
		c.Assert(access.Enable, Equals, true)
	})
}
