package parser

import (
	"path/filepath"

	"github.com/slowmanchan/jsonToCsv/jsonparser"
	"github.com/slowmanchan/jsonToCsv/xmlparser"
)

type Parser interface {
	Read(data []byte) error
	Write() error
}

func New(fileName, outFile, headerFile string) Parser {
	ext := filepath.Ext(fileName)
	switch ext {
	case ".json":
		return jsonparser.New(fileName, outFile, headerFile)
	case ".xml":
		return xmlparser.New(fileName, outFile, headerFile)
	default:
		return nil
	}
}
