package models

type (
	// Movie represents movie object from omdb response
	Movie struct {
		Title      string  `json:"title"`
		Year       string  `json:"year"`
		ImdbID     string  `json:"imdbID"`
		Type       string  `json:"type"`
		Poster     string  `json:"poster"`
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

	// MoviePayloadList represents query params to fetch all movies
	MoviePayloadList struct {
		Pagination int    `query:"pagination" validate:"required"`
		Searchword string `query:"searchword" validate:"required"`
	}

	// MoviePayloadDetail represents query params to get movie's detail
	MoviePayloadDetail struct {
		Id string `param:"id" validate:"required"`
	}
)
