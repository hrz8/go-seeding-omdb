package error

import "errors"

type (
	errorMap struct {
		Status int
		Err    error
	}
)

var (
	List = errorMap{
		Status: 400,
		Err:    errors.New("failed to fetch movies list"),
	}
	Detail = errorMap{
		Status: 400,
		Err:    errors.New("failed to fetch movie detail"),
	}
	DetailNotFound = errorMap{
		Status: 404,
		Err:    errors.New("imdb id not found"),
	}
	ListNotFound = errorMap{
		Status: 404,
		Err:    errors.New("searchword doesn't match with any of available data"),
	}
)
