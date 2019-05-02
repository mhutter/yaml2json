package convert_test

import (
	"bytes"
	"testing"

	"github.com/mhutter/yaml2json/convert"
	"github.com/stretchr/testify/assert"
)

func TestConvert(t *testing.T) {
	in := map[interface{}]interface{}{
		"map": map[interface{}]interface{}{
			"foo": "bar",
		},
		2:    "two",
		true: true,
	}
	expected := map[string]interface{}{
		"map": map[string]interface{}{
			"foo": "bar",
		},
		"2":    "two",
		"true": true,
	}

	out := convert.Map(in)
	assert.Equal(t, expected, out)
}

type tc struct {
	in string
	pretty bool
	out string
}
var cases = []tc{
	{
		in: `foo bar`,
		pretty: false,
		out: `"foo bar"
`,
	},
	{
		in: `---
foo: &dad
  name: Foo
  hoor: viel
  schnaebi: true

bar:
  <<: *dad
  name: Bar
`,
		pretty: false,
		out: `{"bar":{"hoor":"viel","name":"Bar","schnaebi":true},"foo":{"hoor":"viel","name":"Foo","schnaebi":true}}
`,
	},
	{
		in: `---
foo: &dad
  name: Foo
  hoor: viel
  schnaebi: true

bar:
  <<: *dad
  name: Bar
`,
		pretty: true,
		out: `{
  "bar": {
    "hoor": "viel",
    "name": "Bar",
    "schnaebi": true
  },
  "foo": {
    "hoor": "viel",
    "name": "Foo",
    "schnaebi": true
  }
}
`,
	},
	{
		in: `---
- foo
- bar
`,
		pretty: false,
		out: `["foo","bar"]
`,
	},
	{
		in: `---
- foo
- bar
`,
		pretty: true,
		out: `[
  "foo",
  "bar"
]
`,
	},
}

func TestYAML2JSON(t *testing.T) {
	for _, c := range cases {
		var out bytes.Buffer
		err := convert.YAML2JSON(bytes.NewBufferString(c.in), &out, c.pretty)
		assert.Nil(t, err)
		assert.Equal(t, c.out, out.String())
	}
}

func TestY2JError(t *testing.T) {
	err := convert.YAML2JSON(bytes.NewBufferString(""), nil, false)
	assert.Error(t, err)
}
