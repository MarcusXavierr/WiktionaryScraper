package scraper_test

import (
	"reflect"
	"testing"

	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
)

func TestParseResponse(t *testing.T) {
	got := scraper.ParseApiResponse(dumb_down_definition)

	exampleOne := []scraper.Example{"Can you believe that scumbag Steve asked to sleep with her before even asking her name?"}
	definitionTwo := scraper.Definition{"A sleazy, disreputable or despicable person; lowlife.", exampleOne}
	want := scraper.Response{
		[]scraper.Usage{
			{PartOfSpeech: "Noun", Language: "English", Definitions: []scraper.Definition{
				{"A condom.", []scraper.Example{}},
				definitionTwo,
			}},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v", want, got)
	}
}

const dumb_down_definition = `{"en":[{"partOfSpeech":"Noun","language":"English","definitions":[{"definition":"A <a rel=\"mw:WikiLink\" href=\"/wiki/condom\" title=\"condom\">condom</a>."},{"definition":"A <a rel=\"mw:WikiLink\" href=\"/wiki/sleazy\" title=\"sleazy\">sleazy</a>, <a rel=\"mw:WikiLink\" href=\"/wiki/disreputable\" title=\"disreputable\">disreputable</a> or <a rel=\"mw:WikiLink\" href=\"/wiki/despicable\" title=\"despicable\">despicable</a> person; <a rel=\"mw:WikiLink\" href=\"/wiki/lowlife\" title=\"lowlife\">lowlife</a>.","parsedExamples":[{"example":"Can you believe that <b>scumbag</b> Steve asked to sleep with her before even asking her name?"}],"examples":["Can you believe that <b>scumbag</b> Steve asked to sleep with her before even asking her name?"]}]}]}`
