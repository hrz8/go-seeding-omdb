package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hrz8/go-seeding-omdb/models"
)

type (
	movie struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	}

	omdbResponse struct {
		Search       *[]movie `json:"Search"`
		TotalResults string   `json:"totalResults"`
		Response     string   `json:"Response"`
	}
)

func FetchOmdb(apiKey *string, payload *models.MoviePayloadList) (*omdbResponse, error) {
	URL := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&s=%s&page=%d",
		*apiKey, payload.Searchword, payload.Pagination,
	)
	var client = &http.Client{}
	var data omdbResponse

	request, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		return nil, err
	}

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
