package app

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/slowmanchan/jsonToCsv/converter"
	"github.com/slowmanchan/jsonToCsv/helpers"
)

type Apper interface {
	Write(data []map[string]interface{}) error
}

type App struct {
	InFileName     string
	OutFileName    string
	HeaderFileName string
	Converter      converter.Converter
}

func New() (*App, error) {
	a := new(App)

	a.setFlags()

	if a.InFileName == "" {
		return nil, errors.New("input file cannot be blank")
	}

	if a.OutFileName == "" {
		ext := filepath.Ext(a.InFileName)
		fmt.Printf("No out file provided, parsing from InFileName %s\n", a.InFileName)
		a.OutFileName = a.InFileName[0:len(a.InFileName)-len(ext)] + ".csv"
	}

	fmt.Printf("Converting %s to csv\n", a.InFileName)
	fmt.Printf("Using outfile %s\n", a.OutFileName)

	c, err := converter.New(a.InFileName)
	if err != nil {
		return nil, err
	}

	a.Converter = c
	return a, nil
}

func (a *App) Write(data []map[string]interface{}) error {
	headers, err := makeHeaders(data, a.HeaderFileName)
	if err != nil {
		return err
	}
	f, err := os.Create(a.OutFileName)
	if err != nil {
		return err
	}
	csvEnc := csv.NewWriter(f)
	if err := csvEnc.Write(headers); err != nil {
		return err
	}

	for _, r := range data {
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

func (a *App) setFlags() {
	flag.StringVar(&a.InFileName, "i", "", "input file")
	flag.StringVar(&a.InFileName, "o", "", "out file")
	flag.StringVar(&a.InFileName, "h", "", "header file")
	flag.Parse()
}

func makeHeaders(rr []map[string]interface{}, h string) ([]string, error) {
	if h == "" {
		fmt.Println("No header file provided. Using headers from input file")
		headers := []string{}
		for _, r := range rr {
			for k := range r {
				if !helpers.Contains(k, headers) {
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
		return nil, err
	}
	enc := csv.NewReader(f)
	headers, err := enc.Read()
	if err != nil {
		return nil, err
	}
	sort.Strings(headers)
	return headers, nil
}
