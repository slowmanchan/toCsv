package jsonconverter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestJsonCoverter(t *testing.T) {
	var tests = []struct {
		testName string
		input    string
		want     []map[string]interface{}
	}{
		{
			input: `[
				{
					"foo":"bar",
					"float": 2.1,
					"array": ["hello", "1"],
					"nested": {
						"n": "1",
						"x": "2"
					}
				}
			]`,
			want: []map[string]interface{}{
				{
					"array.0":  "hello",
					"array.1":  "1",
					"foo":      "bar",
					"float":    2.1,
					"nested.n": "1",
					"nested.x": "2",
				},
			},
		},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%s", tt.testName)
		t.Run(testName, func(t *testing.T) {
			j := New([]byte(tt.input))
			res, err := j.Convert()
			fmt.Println(err)
			eq := reflect.DeepEqual(res, tt.want)
			if !eq {
				t.Errorf("got %v, want %v", res, tt.want)
			}
		})
	}
}
