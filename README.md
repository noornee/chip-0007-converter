## WIP ⚠️
# chip-0007-converter

## how it works
execute  `go run *.go`
<i> this would read the default csv file `sample.csv` </i>

### NOTE:
<p> if you want to use an external csv, the column header of the csv file should be of the same type as the one provided in the sample </p>

i.e.
`Series Number,File Name,Description,Gender,UUID`


execute 
`go run *.go --help`
in the parent directory to receive a guide on how use the flags.


## using flags
`go run *.go --csv <path to csv file>`

i.e. `go run . --csv file.csv`

this would read the csv file and generate a json file containing the parsed data from the csv in the format of `chip-0007`
