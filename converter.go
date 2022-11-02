package main

type MetaData struct {
	Format           string     `json:"format"`
	Name             string     `json:"name"`
	Description      string     `json:"description"`
	MintingTool      string     `json:"minting_tool"`
	SensitiveContent bool       `json:"sensitive_content"`
	SeriesNumber     int        `json:"series_number"`
	SeriesTotal      int        `json:"series_total"`
	Attributes       []attr1    `json:"attributes"`
	Collection       collection `json:"collection"`
	Data             data       `json:"data"`
	Hash             string     `json:"hash"`
}

type attr1 struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	MinValue  int    `json:"min_value,omitempty"`
	MaxValue  int    `json:"max_value,omitempty"`
}

type attr2 struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type collection struct {
	Name       string  `json:"name"`
	ID         string  `json:"id"`
	Attributes []attr2 `json:"attributes"`
}

type data struct {
	ExampleData string `json:"example_data"`
}
