package records

import (
	"context"
	"nirikshan-backend/pkg/entities"
	"nirikshan-backend/pkg/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Repository holds the mongo database implementation of the Service
type Repository interface {
	GetDump(siteName string) (*[]entities.UserRecords, error)
	CreateDump(configs *entities.UserRecords) error
}

type repository struct {
	Collection *mongo.Collection
}

func (r repository) GetDump(siteName string) (*[]entities.UserRecords, error) {
	var dumpRecords []entities.UserRecords
	cursor, err := r.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	for cursor.Next(context.TODO()) {
		var dumpRecord entities.UserRecords
		_ = cursor.Decode(&dumpRecord)
		dumpRecords = append(dumpRecords, dumpRecord)
	}
	return &dumpRecords, nil
}

func (r repository) CreateDump(configs *entities.UserRecords) error {
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
