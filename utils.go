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
