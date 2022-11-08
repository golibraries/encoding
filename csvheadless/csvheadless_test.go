package csvheadless_test

import (
	"testing"

	"github.com/golibraries/encoding/csvheadless"

	. "github.com/frankban/quicktest"
)

type Access struct {
	Type    int64  `csv:"type"`
	Account []byte `csv:"account"`
	Enable  bool   `csv:"enable"`
}

func TestCSV(t *testing.T) {
	c := New(t)
	c.Run("TestCSV", func(c *C) {
		data, err := csvheadless.Marshal([]*Access{{
			Type:    1,
			Account: []byte("123456"),
			Enable:  true,
		}})
		c.Assert(err, IsNil)

		accesses := []*Access{}
		err = csvheadless.Unmarshal(data, &accesses)
		c.Assert(err, IsNil)
		c.Assert(accesses[0].Type, Equals, int64(1))
		c.Assert(accesses[0].Account, DeepEquals, []byte("123456"))
		c.Assert(accesses[0].Enable, Equals, true)
	})
}
