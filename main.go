package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/slowmanchan/jsonToCsv/parser"
)

func main() {
	inFile := flag.String("i", "", "input file")
	outFile := flag.String("o", "", "output file")
	headerFile := flag.String("h", "", "header file")
	flag.Parse()

	fmt.Printf("Converting %s to csv\n", *inFile)

	data, err := ioutil.ReadFile(*inFile)
	if err != nil {
		log.Fatal(err)
	}

	p := parser.New(*inFile, *outFile, *headerFile)
	err = p.Read(data)
	if err != nil {
		log.Fatal(err)
	}

	err = p.Write()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Finished Converting to csv")
}
