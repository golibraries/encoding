package yaml_test

import (
	"testing"

	"github.com/golibraries/encoding/yaml"

	. "github.com/frankban/quicktest"
)

type Access struct {
	Type    int64  `yaml:"type"`
	Account string `yaml:"account"`
	Enable  bool   `yaml:"enable"`
}

func TestYaml(t *testing.T) {
	c := New(t)
	c.Run("TestYaml", func(c *C) {
		data, err := yaml.Marshal(&Access{
			Type:    1,
			Account: "123456",
			Enable:  true,
		})
		c.Assert(err, IsNil)
		c.Assert(string(data), Equals, `type: 1
account: "123456"
enable: true
`)

		access := &Access{}
		err = yaml.Unmarshal(data, access)
		c.Assert(err, IsNil)
		c.Assert(access.Type, Equals, int64(1))
		c.Assert(access.Account, DeepEquals, "123456")
		c.Assert(access.Enable, Equals, true)
	})
}
