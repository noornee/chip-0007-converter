package main

type CHIP struct {
	Format           string  `json:"format"`
	Name             string  `json:"name"`
	Description      string  `json:"description"`
	MintingTool      string  `json:"minting_tool"`
	SensitiveContent bool    `json:"sensitive_content"`
	SeriesNumber     int     `json:"series_number"`
	SeriesTotal      int     `json:"series_total"`
	Attributes       []Attr1 `json:"attributes"`
	Collection       Collect `json:"collection"`
	Data             Dat     `json:"data"`
}

type Attr1 struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
	MinValue  int    `json:"min_value,omitempty"`
	MaxValue  int    `json:"max_value,omitempty"`
}

type Attr2 struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Collect struct {
	Name       string  `json:"name"`
	ID         string  `json:"id"`
	Attributes []Attr2 `json:"attributes"`
}

type Dat struct {
	ExampleData string `json:"example_data"`
}
