package scraper_test

import (
	"reflect"
	"testing"

	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
)

func TestNormalizeResponse(t *testing.T) {
	response := mockResponseStruct()

	got := response.Normalize()

	example := []scraper.Example{"Can you believe that scumbag Steve asked to sleep with her before even asking her name?"}
	definition := scraper.Definition{"A sleazy, disreputable or despicable person; lowlife.", example}
	want := scraper.Response{
		[]scraper.Usage{
			{PartOfSpeech: "Noun", Language: "English", Definitions: []scraper.Definition{
				definition,
			}},
		},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("expected %v but got %v", want, got)
	}
}
