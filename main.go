package main

func main() {
	dataSchema := &CHIP{
		Format:           "",
		Name:             "",
		Description:      "",
		MintingTool:      "",
		SensitiveContent: false,
		SeriesNumber:     0,
		SeriesTotal:      0,
		Attributes:       []Attr1{},
		Collection:       Collect{},
		Data:             Dat{},
	}
}
