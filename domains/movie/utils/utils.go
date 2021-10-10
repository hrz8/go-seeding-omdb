package utils

import (
	"encoding/json"
	"fmt"

	"github.com/hrz8/go-seeding-omdb/helpers"
	"github.com/hrz8/go-seeding-omdb/models"
)

type (
	movie struct {
		Title      string  `json:"Title"`
		Year       string  `json:"Year"`
		ImdbID     string  `json:"imdbID"`
		Type       string  `json:"Type"`
		Poster     string  `json:"Poster"`
		Released   *string `json:"released,omitempty"`
		Runtime    *string `json:"runtime,omitempty"`
		Director   *string `json:"director,omitempty"`
		Writer     *string `json:"writer,omitempty"`
		Actors     *string `json:"actors,omitempty"`
		Plot       *string `json:"plot,omitempty"`
		Language   *string `json:"language,omitempty"`
		Country    *string `json:"country,omitempty"`
		ImdbRating *string `json:"imdbRating,omitempty"`
		ImdbVotes  *string `json:"imdbVote,omitempty"`
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
	var data omdbResponse

	response, err := helpers.Fetch("GET", URL)
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
