package goenvy

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Getenvs(t *testing.T) {
	// want := "Goenvy"
	// if got := Goenvy("test.env.json"); got != want {
	// t.Errorf("Goenvy() = %q, want %q", got, want)
	// }
}

func Test_toPaths(t *testing.T) {
	raw := `
{
	"string": "string",
	"number": 1,
	"bool": true,
	"object": {
		"nestedKey": "nestedValue"
	},
	"array": [
		{
			"objectInArrayKey1": "objectInArrayValue1"
		},
		{
			"objectInArrayKey2": "objectInArrayValue2"
		}
	]
}
	`
	expected := map[string]string{
		"STRING":                    "string",
		"NUMBER":                    "1",
		"BOOL":                      "true",
		"OBJECT_NESTEDKEY":          "nestedValue",
		"ARRAY_0_OBJECTINARRAYKEY1": "objectInArrayValue1",
		"ARRAY_1_OBJECTINARRAYKEY2": "objectInArrayValue2",
	}

	var (
		parsed interface{}
		res    = make(map[string]string)
	)
	assert.Nil(t, json.Unmarshal([]byte(raw), &parsed))
	toPaths(parsed, &res, "_", nil)
	assert.True(t, len(res) == len(expected))
	for k, v := range expected {
		assert.Equal(t, v, res[k])
	}

}
