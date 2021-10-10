package models

type (
	// Movie represents movie object from omdb response
	Movie struct {
		Title  string `json:"title"`
		Year   string `json:"year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"type"`
		Poster string `json:"poster"`
	}

	// MoviePayloadList represents query params to fetch all movies
	MoviePayloadList struct {
		Pagination int    `query:"pagination"`
		Searchword string `query:"searchword"`
	}

	// MoviePayloadDetail represents query params to get movie's detail
	MoviePayloadDetail struct {
		Id string `param:"id" validate:"required"`
	}
)
