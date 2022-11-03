package main

import (
	"flag"
	"os"
)

func main() {

	csvFilePath := flag.String("csv", "file.csv", "path to the csv file")
	flag.Parse()

	file, err := os.Open(*csvFilePath) // open csv file
	check(err)
	defer file.Close()

	convertCSVtoJSON(file)
	convertJSONtoCSV("output.json", "filename.output.csv")

}
