package api

import "testing"

func TestParseUrlCorrectly(t *testing.T) {
	baseUrl := "https://en.wiktionary.org/api/rest_v1/page/definition/"
	word := "More often than not"
	got := mountUrl(baseUrl, word)
	want := baseUrl + "more_often_than_not"

	if got != want {
		t.Errorf("expected %s but got %s", got, want)
	}
}
