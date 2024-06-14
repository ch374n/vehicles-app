//go:generate mockgen -source manufacturer_repo.go -destination mock/manufacturer_repo_mock.go -package mock MockCollection
package repository

import (
	"context"
	"github.com/ch374n/vehicles-app/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManufacturerRepo interface {
	GetAllManufacturers(ctx context.Context, coll *mongo.Collection) ([]models.Manufacturer, error)
	GetManufacturer(ctx context.Context, coll *mongo.Collection, id int) (models.Manufacturer, error)
	CreateManufacturer(ctx context.Context, coll *mongo.Collection, manufacturer models.Manufacturer) error
	UpdateManufacturer(ctx context.Context, coll *mongo.Collection, id int, manufacturer models.Manufacturer) error
	DeleteManufacturer(ctx context.Context, coll *mongo.Collection, id int) error
}

type ManufacturerRepoImpl struct {
}

func NewManufacturerRepo() ManufacturerRepo {
	return &ManufacturerRepoImpl{}
}

func (r *ManufacturerRepoImpl) GetAllManufacturers(ctx context.Context, coll *mongo.Collection) ([]models.Manufacturer, error) {
	cursor, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var manufacturers []models.Manufacturer
	if err = cursor.All(ctx, &manufacturers); err != nil {
		return nil, err
	}

	return manufacturers, nil
}

func (r *ManufacturerRepoImpl) GetManufacturer(ctx context.Context, coll *mongo.Collection, id int) (models.Manufacturer, error) {
	var manufacturer models.Manufacturer
	err := coll.FindOne(ctx, bson.M{"Mfr_ID": id}).Decode(&manufacturer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return manufacturer, nil
		}
		return manufacturer, err
	}

	return manufacturer, nil
}

func (r *ManufacturerRepoImpl) CreateManufacturer(ctx context.Context, coll *mongo.Collection, manufacturer models.Manufacturer) error {
	_, err := coll.InsertOne(ctx, manufacturer)
	return err
}

func (r *ManufacturerRepoImpl) UpdateManufacturer(ctx context.Context, coll *mongo.Collection, id int, manufacturer models.Manufacturer) error {
	_, err := coll.UpdateOne(ctx, bson.M{"Mfr_ID": id}, bson.M{"$set": manufacturer})
	return err
}

func (r *ManufacturerRepoImpl) DeleteManufacturer(ctx context.Context, coll *mongo.Collection, id int) error {
	_, err := coll.DeleteOne(ctx, bson.M{"Mfr_ID": id})
	return err
}
