package siteconfigs

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"nirikshan-backend/pkg/entities"
	"nirikshan-backend/pkg/utils"
)

// Repository holds the mongo database implementation of the Service
type Repository interface {
	GetSiteData(siteName string) (*entities.SiteConfigs, error)
	CreateSiteData(configs *entities.SiteConfigs) error
}

type repository struct {
	Collection *mongo.Collection
}

func (r repository) GetSiteData(siteName string) (*entities.SiteConfigs, error) {
	var result = entities.SiteConfigs{}
	findOneErr := r.Collection.FindOne(context.TODO(), bson.M{
		"site_name": siteName,
	}).Decode(&result)

	if findOneErr != nil {
		return nil, findOneErr
	}
	return &result, nil
}

func (r repository) CreateSiteData(configs *entities.SiteConfigs) error {
	_, err := r.Collection.InsertOne(context.Background(), configs)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return utils.ErrUserExists
		}
		return err
	}
	return nil
}

// NewRepo creates a new instance of this repository
func NewRepo(collection *mongo.Collection) Repository {
	return &repository{
		Collection: collection,
	}
}
