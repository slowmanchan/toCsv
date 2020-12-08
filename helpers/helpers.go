package helpers

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/jeremywohl/flatten"
	"github.com/pkg/errors"
)

func WriteMapToCSV(results []map[string]interface{}, outFile, headerFile string) error {
	headers, err := makeHeaders(results, headerFile)
	if err != nil {
		return err
	}
	f, err := os.Create(outFile)
	if err != nil {
		return err
	}
	csvEnc := csv.NewWriter(f)
	if err := csvEnc.Write(headers); err != nil {
		return err
	}

	for _, r := range results {
		values := []string{}
		for _, h := range headers {
			s, _ := r[h].(string)
			s = strings.Replace(s, ",", " ", -1)
			s = strings.Replace(s, "\n", " ", -1)
			s = strings.Replace(s, "\r", " ", -1)
			values = append(values, s)
		}
		if err := csvEnc.Write(values); err != nil {
			return err
		}
	}
	return nil
}

func SetOutFile(inFile, outFile string) string {
	if outFile == "" {
		ext := filepath.Ext(inFile)
		fmt.Printf("No out file provided, parsing from inFile %s\n", outFile)
		return inFile[0:len(inFile)-len(ext)] + ".csv"
	}
	fmt.Printf("Using outfile %s\n", outFile)
	return outFile
}

func makeHeaders(rr []map[string]interface{}, h string) ([]string, error) {
	if h == "" {
		fmt.Println("No header file provided. Using headers from input file")
		headers := []string{}
		for _, r := range rr {
			for k := range r {
				if !contains(k, headers) {
					headers = append(headers, k)
				}
			}
		}
		sort.Strings(headers)
		return headers, nil
	}

	fmt.Printf("Using header file %s\n", h)
	f, err := os.Open(h)
	if err != nil {
		log.Fatal(err)
	}
	enc := csv.NewReader(f)
	headers, err := enc.Read()
	if err != nil {
		log.Fatal(err)
	}
	sort.Strings(headers)
	return headers, nil
}

func contains(k string, ss []string) bool {
	for _, s := range ss {
		if k == s {
			return true
		}
	}
	return false
}

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
