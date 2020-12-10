package xmlconverter

import (
	"fmt"
	"reflect"
	"testing"
)

func TestXMLConverter(t *testing.T) {
	var tests = []struct {
		testName string
		input    string
		want     []map[string]interface{}
	}{
		{
			input: `
			<?xml version="1.0" encoding="UTF-8"?>
			<rows>
			<row>
				<foo>bar</foo>
			</row>
			</rows>
			`,
			want: []map[string]interface{}{
				{
					"foo": "bar",
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
