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

		series_number, _ := strconv.Atoi(line[0])

		data := &MetaData{
			Format:           "CHIP-0007",
			Name:             line[1],
			Description:      line[2],
			MintingTool:      "",
			SensitiveContent: false,
			SeriesNumber:     series_number,
			SeriesTotal:      0,
			Attributes: []attr1{{
				TraitType: "Gender",
				Value:     line[3],
			}},
			Collection: collection{
				Name: line[1],
				ID:   line[4],
				Attributes: []attr2{{
					Type:  "",
					Value: "",
				}},
			},
			Data: data{
				ExampleData: "",
			},
			//Hash: line[5],
		}

		// convert struct to []bytes to be able to hash it
		dataByte, err := json.Marshal(data)
		check(err)
		data.Hash = hash256(dataByte) // update struct with the value of the hash

		records = append(records, *data)
	}

	jsonData, err := json.MarshalIndent(records, "", " ")
	check(err)

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

	header := []string{"SeriesNumber", "File Name", "Description", "Gender", "UUID", "Hash"} // header for csv file
	w.Write(header)

	for _, obj := range d {
		var record []string
		//record = append(record, obj.Name, obj.Description, obj.Hash)
		record = append(record, fmt.Sprintf("%d", obj.SeriesNumber), obj.Name, obj.Description, "" ,obj.Collection.ID, obj.Hash)
		w.Write(record)
	}
}
