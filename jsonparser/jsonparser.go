package jsonparser

import (
	"encoding/json"

	"github.com/slowmanchan/jsonToCsv/helpers"
)

type JsonParser struct {
	inFile     string
	outFile    string
	headerFile string
	data       []map[string]interface{}
}

func New(inFile, outFile, headerFile string) *JsonParser {
	outFile = helpers.SetOutFile(inFile, outFile)
	return &JsonParser{
		headerFile: headerFile,
		inFile:     inFile,
		outFile:    outFile,
	}
}

func (j *JsonParser) Read(data []byte) error {
	result := []map[string]interface{}{}
	if err := json.Unmarshal(data, &result); err != nil {
		return err
	}
	f, err := helpers.FlattenMaps(result)
	if err != nil {
		return err
	}
	j.data = f
	return nil
}

func (j *JsonParser) Write() error {
	return helpers.WriteMapToCSV(j.data, j.outFile, j.headerFile)
}
