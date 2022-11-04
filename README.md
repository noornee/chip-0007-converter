
## WIP ⚠️

Fri Nov  4 11:32:22 WAT 2022
theres an issue reverting back to csv

# chip-0007-converter

## how it works
execute  `go run *.go`
<i> this would read the default csv file `sample.csv` </i>
this generates 2 files: `output.json` and `filename.output.csv`

### NOTE:
<p> if you want to use an external csv, the column header of the csv file should be of the same type as the one provided in the sample </p>

i.e.
`TEAM NAMES,Series Number,Filename,Name,Description,Gender,Attributes,UUID`


execute 
`go run *.go --help`
in the parent directory to receive a guide on how use the flags.


## using flags
`go run *.go --csv <path to csv file>`

i.e. `go run *.go --csv file.csv`

this would read the csv file and generate a json file containing the parsed data from the csv in the format of `chip-0007`


## image 
![Screenshot_20221104-125019](https://user-images.githubusercontent.com/71889751/199966046-19521c89-a7f9-4365-90fd-b7db2b13d725.png)
