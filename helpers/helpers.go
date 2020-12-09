package helpers

import (
	"log"

	"github.com/jeremywohl/flatten"
	"github.com/pkg/errors"
)

func FlattenMaps(rows []map[string]interface{}) ([]map[string]interface{}, error) {
	for i, row := range rows {
		flatRow, err := flatten.Flatten(row, "", flatten.DotStyle)
		if err != nil {
			return nil, errors.Wrap(err, "failed to flatten map")
		}
		rows[i] = flatRow
	}
	return rows, nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Contains(k string, ss []string) bool {
	for _, s := range ss {
		if k == s {
			return true
		}
	}
	return false
}
