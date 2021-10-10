package utils

import (
	"encoding/json"
	"fmt"

	LogRequestUsecase "github.com/hrz8/go-seeding-omdb/domains/log_request/usecase"
	MovieError "github.com/hrz8/go-seeding-omdb/domains/movie/error"
	"github.com/hrz8/go-seeding-omdb/helpers"
	"github.com/hrz8/go-seeding-omdb/models"
	Utils "github.com/hrz8/go-seeding-omdb/utils"
)

type (
	omdbDetailResponse struct {
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
		Response   string  `json:"Response"`
	}

	omdbListResponse struct {
		Search       *[]omdbDetailResponse `json:"Search"`
		TotalResults string                `json:"totalResults"`
		Response     string                `json:"Response"`
	}
)

func FetchOmdbDetail(apiKey *string, id *string) (*omdbDetailResponse, error) {
	URL := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&i=%s",
		*apiKey, *id,
	)
	var data omdbDetailResponse

	response, err := helpers.Fetch("GET", URL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Response == "False" {
		return nil, MovieError.DetailNotFound.Err
	}

	return &data, nil
}

func FetchOmdbList(apiKey *string, payload *models.MoviePayloadList) (*omdbListResponse, error) {
	URL := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&s=%s&page=%d",
		*apiKey, payload.Searchword, payload.Pagination,
	)
	var data omdbListResponse

	response, err := helpers.Fetch("GET", URL)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Response == "False" {
		return nil, MovieError.ListNotFound.Err
	}

	return &data, nil
}

func LoggingToDB(ctx *Utils.CustomContext, u LogRequestUsecase.UsecaseInterface, typ string) {
	requestUri := ctx.Context.Request().RequestURI
	method := ctx.Context.Request().Method
	headerTemp, err := json.Marshal(ctx.Context.Request().Header)
	if err != nil {
		fmt.Println(err)
	}
	header := string(headerTemp)
	u.Create(ctx, &models.LogRequest{
		Type:       typ,
		RequestURI: requestUri,
		Header:     header,
		Method:     method,
	})
}
