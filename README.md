This is a command line tool that will convert a JSON or XML file to a CSV file.

## Usage

### Example
`go run main.go -i test.json -o test.csv`

### Flags

```
-i file that you want converted
-o (optional) filename of output. will default to the input file name and add .csv
-h (optional) header file. include only the headers you want. will infer and include all by default
```

### Extend

Implement the `Converter` interface to add more data types

