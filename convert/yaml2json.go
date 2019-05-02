package convert

import (
	"encoding/json"
	"fmt"
	"io"

	"gopkg.in/yaml.v3"
)

// YAML2JSON reads YAML data from in, converts it into JSON and writes
// it into out
func YAML2JSON(in io.Reader, out io.Writer, pretty bool) error {
	var data interface{}

	if err := yaml.NewDecoder(in).Decode(&data); err != nil {
		return err
	}
	switch d2 := data.(type) {
	case map[interface{}]interface{}:
		data = Map(d2)
	}

	enc := json.NewEncoder(out)
	if pretty {
		enc.SetIndent("", "  ")
	}
	return enc.Encode(data)
}

// Map converts all keys in a map to strings (as required by
// encoding/json).
func Map(m map[interface{}]interface{}) map[string]interface{} {
	res := map[string]interface{}{}

	for k, v := range m {
		switch v2 := v.(type) {
		case map[interface{}]interface{}:
			res[fmt.Sprint(k)] = Map(v2)
		default:
			res[fmt.Sprint(k)] = v2
		}
	}

	return res
}
