package repository

import (
	"strconv"

	"github.com/hrz8/go-seeding-omdb/domains/movie/utils"
	"github.com/hrz8/go-seeding-omdb/models"
)

type (
	RepositoryInterface interface {
		List(apiKey *string, payload *models.MoviePayloadList) (*[]models.Movie, *int, error)
		Detail(apiKey *string, id *string) (*models.Movie, error)
	}

	impl struct{}
)

func (i *impl) List(apiKey *string, payload *models.MoviePayloadList) (*[]models.Movie, *int, error) {
	response, err := utils.FetchOmdbList(apiKey, payload)
	if err != nil {

	}
	total, err := strconv.Atoi(response.TotalResults)
	if err != nil {

	}
	result := make([]models.Movie, len(*response.Search))
	for index, movie := range *response.Search {
		result[index] = models.Movie{
			Title:  movie.Title,
			Year:   movie.Year,
			ImdbID: movie.ImdbID,
			Type:   movie.Type,
			Poster: movie.Poster,
		}
	}
	return &result, &total, nil
}

func (i *impl) Detail(apiKey *string, id *string) (*models.Movie, error) {
	response, err := utils.FetchOmdbDetail(apiKey, id)
	if err != nil {

	}
	result := &models.Movie{
		Title:      response.Title,
		Year:       response.Year,
		ImdbID:     response.ImdbID,
		Type:       response.Type,
		Poster:     response.Poster,
		Released:   response.Released,
		Runtime:    response.Runtime,
		Director:   response.Director,
		Writer:     response.Writer,
		Actors:     response.Actors,
		Plot:       response.Plot,
		Language:   response.Language,
		Country:    response.Country,
		ImdbRating: response.ImdbRating,
		ImdbVotes:  response.ImdbVotes,
	}
	return result, nil
}

func NewRepository() RepositoryInterface {
	return &impl{}
}
