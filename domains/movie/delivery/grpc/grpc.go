package grpc

import (
	"context"

	"github.com/hrz8/go-seeding-omdb/models"
	"google.golang.org/grpc"
)

type (
	MoviesServer struct {
		models.UnimplementedMoviesServer
	}
)

func (i MoviesServer) List(ctx context.Context, payload *models.MoviePayloadList) (*models.MovieList, error) {
	return &models.MovieList{}, nil
}

func (i MoviesServer) Detail(ctx context.Context, payload *models.MoviePayloadDetail) (*models.Movie, error) {
	return &models.Movie{}, nil
}

func RegisterServer(s *grpc.Server, gServer models.MoviesServer) {
	models.RegisterMoviesServer(s, gServer)
}
