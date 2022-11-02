package main

type CHIP struct {
	Format           string     `csv:"format"`
	Name             string     `csv:"name"`
	Description      string     `csv:"description"`
	MintingTool      string     `csv:"minting_tool"`
	SensitiveContent bool       `csv:"sensitive_content"`
	SeriesNumber     int        `csv:"series_number"`
	SeriesTotal      int        `csv:"series_total"`
	Attributes       []attr1    `csv:"attributes"`
	Collection       collection `csv:"collection"`
	Data             data       `csv:"data"`
}

type attr1 struct {
	TraitType string `csv:"trait_type"`
	Value     string `csv:"value"`
	MinValue  int    `csv:"min_value,omitempty"`
	MaxValue  int    `csv:"max_value,omitempty"`
}

type attr2 struct {
	Type  string `csv:"type"`
	Value string `csv:"value"`
}

type collection struct {
	Name       string  `csv:"name"`
	ID         string  `csv:"id"`
	Attributes []attr2 `csv:"attributes"`
}

type data struct {
	ExampleData string `csv:"example_data"`
}
