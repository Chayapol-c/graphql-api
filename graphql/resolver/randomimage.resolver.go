package resolvers

import (
	"github.com/graphql-go/graphql"
	"graphql-tutorial/data/models"
	"graphql-tutorial/internal/repository"
)

func GetImageResolve(params graphql.ResolveParams) (interface{}, error) {

	// Fetch Image from rest api
	image_url, err := repository.RandomImageUrl(150, 150)
	image := models.ImageModel{ImageUrl: image_url}
	if err != nil {
		return nil, err
	}
	return image, nil
}
