package adapter

import (
	"link-shortener/usecases"

	"go.mongodb.org/mongo-driver/mongo"
)

type urlMappingRepo struct {
	database *mongo.Database
}

func NewUrlMappingRepo(database *mongo.Database) usecases.UrlMappingRepo {
	return &urlMappingRepo{database: database}
}

func (r *urlMappingRepo) CreateUrlMapping(string) (string, error) {
	// Implement the logic to create a URL mapping in the repository
	return "", nil
}
