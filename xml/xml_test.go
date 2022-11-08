package xml_test

import (
	"testing"

	"github.com/golibraries/encoding/xml"

	. "github.com/frankban/quicktest"
)

type Access struct {
	Type    int64  `xml:"type"`
	Account []byte `xml:"account"`
	Enable  bool   `xml:"enable"`
}

func TestXML(t *testing.T) {
	c := New(t)
	c.Run("TestXML", func(c *C) {
		data, err := xml.Marshal(&Access{
			Type:    1,
			Account: []byte("123456"),
			Enable:  true,
		})
		c.Assert(err, IsNil)
		c.Assert(string(data), Equals, `<Access><type>1</type><account>123456</account><enable>true</enable></Access>`)

		access := &Access{}
		err = xml.Unmarshal(data, access)
		c.Assert(err, IsNil)
		c.Assert(access.Type, Equals, int64(1))
		c.Assert(access.Enable, Equals, true)
		c.Assert(access.Account, DeepEquals, []byte("123456"))
	})
}
