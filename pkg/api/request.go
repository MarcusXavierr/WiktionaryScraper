package api

import (
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/MarcusXavierr/wiktionary-scraper/pkg/scraper"
	"github.com/ettle/strcase"
)

var NotFoundError = errors.New("Could not get definition")

func GetDefinition(url, sentence string) (scraper.Response, error) {
	resp, err := http.Get(mountUrl(url, sentence))
	if err != nil {
		return scraper.Response{}, err
	}

	if resp.StatusCode != http.StatusOK {
		return scraper.Response{}, NotFoundError
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return scraper.Response{}, err
	}

	return scraper.ParseApiResponse(string(body)), nil
}

func mountUrl(baseUrl, sentence string) string {
	return baseUrl + strcase.ToSnake(sentence)
}
