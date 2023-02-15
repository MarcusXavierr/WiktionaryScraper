package api_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/MarcusXavierr/wiktionary-scraper/pkg/api"
	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
)

func TestGetDefinition(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		server := makeServer(http.StatusOK)

		got, _ := api.GetDefinition(server.URL+"/", "More often than not")

		want := mockParsedResponse

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, but got %v", want, got)
		}
	})

	t.Run("returns error when the definition does not exists", func(t *testing.T) {
		server := makeServer(http.StatusNotFound)

		got, err := api.GetDefinition(server.URL+"/", "whatever")

		if err == nil {
			t.Fatal("expected error")
		}

		want := scraper.Response{}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v, but got %v", want, got)
		}
	})
}

func makeServer(statusHttp int) *httptest.Server {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(statusHttp)
		w.Write([]byte(more_often_than_not_response))
	}))
	return server
}

var mockParsedResponse = scraper.Response{
	Usages: []scraper.Usage{
		{Language: "English", PartOfSpeech: "Adverb", Definitions: []scraper.Definition{mockDefinition}},
	},
}

var mockDefinition = scraper.Definition{
	WordDefinition: "more than half the time; more likely to be the case than not to be the case.",
	Examples: []scraper.Example{
		"More often than not, tomato seeds will sprout even if they are a couple of years old.",
	},
}

const more_often_than_not_response = `{"en":[{"partOfSpeech":"Adverb","language":"English","definitions":[{"definition":"<a rel=\"mw:WikiLink\" href=\"/wiki/usually\" title=\"usually\"></a>more than half the time; more <a rel=\"mw:WikiLink\" href=\"/wiki/likely\" title=\"likely\">likely</a> to be the case than not to be the case.","parsedExamples":[{"example":"<b>More often than not</b>, tomato seeds will sprout even if they are a couple of years old."}],"examples":["<b>More often than not</b>, tomato seeds will sprout even if they are a couple of years old."]}]}]}`
