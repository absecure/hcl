package hcl

import (
	hclParser "github.com/hashicorp/hcl/hcl/parser"
	jsonParser "github.com/hashicorp/hcl/json/parser"
)

func Fuzz(data []byte) int {
	switch lexMode(data) {
	case lexModeHcl:
		_, _ = hclParser.Parse(data)
	case lexModeJson:
		_, _ = jsonParser.Parse(data)
	}
	return 1
}
