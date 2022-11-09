package csv_test

import (
	"testing"

	. "github.com/frankban/quicktest"
	"github.com/golibraries/encoding/csv"
)

func TestCSVHeadless(t *testing.T) {
	c := New(t)
	c.Run("TestCSV", func(c *C) {
		data, err := csv.MarshalHeadless([]*Access{{
			Type:    1,
			Account: "123456",
			Enable:  true,
		}})
		c.Assert(err, IsNil)

		accesses := []*Access{}
		err = csv.UnmarshalHeadless(data, &accesses)
		c.Assert(err, IsNil)
		c.Assert(accesses[0].Type, Equals, int64(1))
		c.Assert(accesses[0].Account, DeepEquals, "123456")
		c.Assert(accesses[0].Enable, Equals, true)
	})
}
