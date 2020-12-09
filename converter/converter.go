package converter

import (
	"errors"
	"io/ioutil"
	"path/filepath"

	"github.com/slowmanchan/jsonToCsv/jsonconverter"
	"github.com/slowmanchan/jsonToCsv/xmlconverter"
)

type Converter interface {
	Convert() ([]map[string]interface{}, error)
}

func New(fileName string) (Converter, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	switch filepath.Ext(fileName) {
	case ".json":
		return jsonconverter.New(data), nil
	case ".xml":
		return xmlconverter.New(data), nil
	default:
		return nil, errors.New("Unrecognized data type, only json and xml supported")
	}
}
