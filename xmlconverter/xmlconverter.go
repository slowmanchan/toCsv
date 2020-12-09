package xmlconverter

import (
	"github.com/clbanning/mxj"
	"github.com/pkg/errors"
	"github.com/slowmanchan/jsonToCsv/helpers"
)

type XmlConverter struct {
	data []byte
}

func New(data []byte) *XmlConverter {
	return &XmlConverter{
		data: data,
	}
}

func (x *XmlConverter) Convert() ([]map[string]interface{}, error) {
	m, err := mxj.NewMapXml(x.data)
	if err != nil {
		return nil, errors.Wrap(err, "unable to parse content as xml")
	}
	dd, err := m.ValuesForPath("*.*")
	if err != nil {
		return nil, err
	}

	res := []map[string]interface{}{}
	for _, d := range dd {
		m, _ := d.(map[string]interface{})
		res = append(res, m)
	}

	return helpers.FlattenMaps(res)
}
