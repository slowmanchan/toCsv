package main

import (
	"fmt"

	"github.com/slowmanchan/jsonToCsv/app"
	"github.com/slowmanchan/jsonToCsv/helpers"
)

func main() {
	a, err := app.New()
	helpers.CheckError(err)

	d, err := a.Converter.Convert()
	helpers.CheckError(err)

	err = a.Write(d)
	helpers.CheckError(err)

	fmt.Println("Finished Converting to csv")
}
