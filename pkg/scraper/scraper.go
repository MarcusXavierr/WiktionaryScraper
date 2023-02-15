package scraper

import (
	"encoding/json"
	"regexp"

	"github.com/samber/lo"
)

type Response struct {
	Usages []Usage `json:"en"`
}

type Usage struct {
	PartOfSpeech string       `json:"partOfSpeech"`
	Language     string       `json:"language"`
	Definitions  []Definition `json:"definitions"`
}

type Definition struct {
	WordDefinition string    `json:"definition"`
	Examples       []Example `json:"examples"`
}

type Example string

func ParseApiResponse(data string) Response {
	response := Response{}
	populateResponse(data, &response)

	cleanResponse(&response)
	return response
}

func populateResponse(data string, response *Response) {
	json.Unmarshal([]byte(data), response)
}

func cleanResponse(response *Response) {
	response.Usages = lo.Map(response.Usages, func(usage Usage, _ int) Usage {
		usage.Definitions = lo.Map(usage.Definitions, cleanDefinition)
		return usage
	})
}

func cleanDefinition(definition Definition, index int) Definition {
	re := regexp.MustCompile("<[^>]*>")
	definition.WordDefinition = re.ReplaceAllString(definition.WordDefinition, "")

	definition.Examples = lo.Map(definition.Examples, func(x Example, _ int) Example {
		return Example(re.ReplaceAllString(string(x), ""))
	})
	return definition
}
