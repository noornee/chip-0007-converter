package main

import (
	"encoding/csv"
	"encoding/json"
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
			Hash: line[5],
		}

		records = append(records, *data)
	}

	jsonData, err := json.MarshalIndent(records, "", " ")
	check(err)

	_ = ioutil.WriteFile("output.json", jsonData, 0644) // write to json file

}

//func json_file_hash256(filename string) {

//file, err := os.Open(filename)
//check(err)
//defer file.Close()

//hash := sha256.New() // Hash in

//if _, err := io.Copy(hash, file); err != nil {
//check(err)
//}

////fmt.Printf("%x", hash.Sum(nil))

//}
