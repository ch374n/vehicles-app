package repository

import (
	"context"

	config "github.com/ch374n/vehicles-app/configs"
	"github.com/ch374n/vehicles-app/internal/models"
	"github.com/ch374n/vehicles-app/pkg/database"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ManufacturerRepo struct {
	coll *mongo.Collection
}

func NewManufacturerRepo(cfg *config.Config) *ManufacturerRepo {
	return &ManufacturerRepo{
		coll: database.MongoClient.Database(cfg.DBName).Collection(cfg.CollectionName),
	}
}

func (r *ManufacturerRepo) GetAllManufacturers() ([]models.Manufacturer, error) {
	cursor, err := r.coll.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var manufacturers []models.Manufacturer
	if err = cursor.All(context.Background(), &manufacturers); err != nil {
		return nil, err
	}

	return manufacturers, nil
}

func (r *ManufacturerRepo) GetManufacturer(id int) (models.Manufacturer, error) {
	var manufacturer models.Manufacturer
	err := r.coll.FindOne(context.Background(), bson.M{"Mfr_ID": id}).Decode(&manufacturer)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return manufacturer, nil 
		}
		return manufacturer, err
	}

	return manufacturer, nil
}

func (r *ManufacturerRepo) CreateManufacturer(manufacturer models.Manufacturer) error {
	_, err := r.coll.InsertOne(context.Background(), manufacturer)
	return err
}

func (r *ManufacturerRepo) UpdateManufacturer(id int, manufacturer models.Manufacturer) error {
	_, err := r.coll.UpdateOne(context.Background(), bson.M{"Mfr_ID": id}, bson.M{"$set": manufacturer})
	return err
}

func (r *ManufacturerRepo) DeleteManufacturer(id int) error {
	_, err := r.coll.DeleteOne(context.Background(), bson.M{"Mfr_ID": id})
	return err
}