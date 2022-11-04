package main

import (
	"crypto/sha256"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

// error checker
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// hash256
func hash256(data []byte) string {
	hash := sha256.New()
	hash.Write([]byte(data))

	sha := hash.Sum(nil)
	return fmt.Sprintf("%x", sha)
}

func convertCSVtoJSON(file *os.File) {

	var records []MetaData // would dump all the metadata here

	// Read CSV file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	check(err)

	for i, line := range lines {
		if i == 0 {
			continue // skip header line
		}

		gender := []attr1{{
			TraitType: "Gender",
			Value:     line[5],
		}}

		attributes := line[6]
		var rec []attr1 // contains the appended attributes
		rec = append(rec, gender...)

		traits := strings.Split(attributes, ";")
		for i := range traits {
			trait := strings.Split(traits[i], ":")
			if len(trait) == 2 {
				attrs := []attr1{{TraitType: strings.Trim(trait[0], " "), Value: strings.Trim(trait[1], " ")}}
				rec = append(rec, attrs...)
			} else {
				attrs := []attr1{{TraitType: trait[0], Value: ""}}
				rec = append(rec, attrs[0])

			}
		}

		series_number, _ := strconv.Atoi(line[1])

		data := &MetaData{
			Format:           "CHIP-0007",
			Name:             line[2],
			Description:      line[4],
			MintingTool:      line[0],
			SensitiveContent: false,
			SeriesNumber:     series_number,
			Attributes:       rec,
			SeriesTotal:      0,
			Collection: collection{
				Name: line[3],
				ID:   line[7],
				Attributes: []attr2{{
					Type:  "",
					Value: "",
				}},
			},
			Data: data{
				ExampleData: "",
			},
		}

		//convert struct to []bytes to be able to hash it
		dataByte, err := json.Marshal(data)
		check(err)
		data.Hash = hash256(dataByte) // update struct with the value of the hash

		records = append(records, *data)
	}

	jsonData, err := json.MarshalIndent(records, "", " ")
	check(err)

	fmt.Println("generating json file")

	_ = ioutil.WriteFile("output.json", jsonData, 0644) // write to json file

}

func convertJSONtoCSV(input, output string) {
	data, err := ioutil.ReadFile(input)
	check(err)

	// Unmarshal JSON data
	var d []MetaData
	err = json.Unmarshal([]byte(data), &d)
	check(err)

	// Create a csv file
	f, err := os.Create(output)
	check(err)
	defer f.Close()

	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)
	defer w.Flush()

	// minting tool is team name

	header := []string{"TeamNames", "SeriesNumber", "FileName", "Name", "Description", "Gender", "Attributes", "UUID", "Hash"} // header for csv file
	w.Write(header)

	fmt.Println("transforming generated json file to csv")

	for _, obj := range d {
		var record []string
		//record = append(record, obj.Name, obj.Description, obj.Hash)
		record = append(record, obj.MintingTool, fmt.Sprintf("%d", obj.SeriesNumber), "", obj.Name, obj.Description, fmt.Sprintf("%v", obj.Attributes), obj.Collection.ID, obj.Hash)
		w.Write(record)
	}
}
