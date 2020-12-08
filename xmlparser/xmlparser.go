package xmlparser

import (
	"github.com/clbanning/mxj"
	"github.com/pkg/errors"
	"github.com/slowmanchan/jsonToCsv/helpers"
)

type XmlParser struct {
	inFile     string
	outFile    string
	headerFile string
	data       []map[string]interface{}
}

func New(inFile, outFile, headerFile string) *XmlParser {
	outFile = helpers.SetOutFile(inFile, outFile)
	return &XmlParser{
		inFile:     inFile,
		outFile:    outFile,
		headerFile: headerFile,
	}
}

func (x *XmlParser) Read(data []byte) error {
	m, err := mxj.NewMapXml(data)
	if err != nil {
		return errors.Wrap(err, "unable to parse content as xml")
	}
	dd, err := m.ValuesForPath("*.*")
	if err != nil {
		return err
	}

	res := []map[string]interface{}{}
	for _, d := range dd {
		m, _ := d.(map[string]interface{})
		res = append(res, m)
	}

	f, err := helpers.FlattenMaps(res)
	if err != nil {
		return err
	}
	x.data = f
	return nil
}

func (x *XmlParser) Write() error {
	return helpers.WriteMapToCSV(x.data, x.outFile, x.headerFile)
}
