package jsonconverter

import (
	"encoding/json"

	"github.com/slowmanchan/jsonToCsv/helpers"
)

type JsonConverter struct {
	data []byte
}

func New(data []byte) *JsonConverter {
	return &JsonConverter{
		data: data,
	}
}

func (j *JsonConverter) Convert() ([]map[string]interface{}, error) {
	result := []map[string]interface{}{}
	if err := json.Unmarshal(j.data, &result); err != nil {
		return nil, err
	}
	return helpers.FlattenMaps(result)
}
