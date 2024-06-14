package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"testing"

	models "github.com/ch374n/vehicles-app/internal/models"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateManufacturer(t *testing.T) {
	manufacturerRepo := NewManufacturerRepo()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	manufacturer := models.Manufacturer{
		Country:       "UNITED STATES (USA)",
		MfrCommonName: "Tesla",
		MfrID:         955,
		MfrName:       "TESLA, INC.",
		VehicleTypes: []models.VehicleType{
			{
				IsPrimary: true,
				Name:      "Passenger Car",
			},
		},
	}

	mt.Run("success", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := manufacturerRepo.CreateManufacturer(context.Background(), manufacturerCollection, manufacturer)
		assert.Nil(t, err)
	})

	mt.Run("failure", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		mt.AddMockResponses(bson.D{{"ok", 0}})

		err := manufacturerRepo.CreateManufacturer(context.Background(), manufacturerCollection, manufacturer)
		assert.Error(t, err)
	})
}

func TestGetManufacturerById(t *testing.T) {
	manufacturerRepo := NewManufacturerRepo()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	manufacturer := models.Manufacturer{
		Country:       "UNITED STATES (USA)",
		MfrCommonName: "Tesla",
		MfrID:         955,
		MfrName:       "TESLA, INC.",
		VehicleTypes: []models.VehicleType{
			{
				IsPrimary: true,
				Name:      "Passenger Car",
			},
		},
	}

	mt.Run("success", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "vehiclesapp.vehicles", mtest.FirstBatch, bson.D{
			{"ok", 1},
			{"country", manufacturer.Country},
			{"mfrCommonName", manufacturer.MfrCommonName},
			{"mfrID", manufacturer.MfrID},
			{"mfrName", manufacturer.MfrName},
			{"vehicleTypes", bson.A{
				bson.D{
					{"isPrimary", manufacturer.VehicleTypes[0].IsPrimary},
					{"name", manufacturer.VehicleTypes[0].Name},
				},
			},
			}},
		))

		manufacturerResponse, err := manufacturerRepo.GetManufacturer(context.Background(), manufacturerCollection, manufacturer.MfrID)
		assert.Nil(t, err)

		assert.Equal(t, manufacturer, manufacturerResponse)

	})

	mt.Run("failure", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		mt.AddMockResponses(bson.D{{"ok", 0}})

		_, err := manufacturerRepo.GetManufacturer(context.Background(), manufacturerCollection, manufacturer.MfrID)
		assert.Error(t, err)
	})

}

func TestDeleteManufacturerById(t *testing.T) {
	manufacturerRepo := NewManufacturerRepo()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll
		mt.AddMockResponses(bson.D{{"ok", 1}, {"acknowledged", true}, {"n", 1}})

		err := manufacturerRepo.DeleteManufacturer(context.Background(), manufacturerCollection, 955)
		assert.Nil(t, err)

	})

	mt.Run("failure", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		mt.AddMockResponses(bson.D{{"ok", 0}, {"acknowledged", true}, {"n", 0}})
		err := manufacturerRepo.DeleteManufacturer(context.Background(), manufacturerCollection, 955)
		assert.NotNil(t, err)
	})

}

func TestUpdateManufacturerById(t *testing.T) {
	manufacturerRepo := NewManufacturerRepo()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	manufacturer := models.Manufacturer{
		Country:       "UNITED STATES (USA)",
		MfrCommonName: "Tesla",
		MfrID:         955,
		MfrName:       "TESLA, INC.",
		VehicleTypes: []models.VehicleType{
			{
				IsPrimary: true,
				Name:      "Passenger Car",
			},
		},
	}

	mt.Run("success", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "vehiclesapp.vehicles", mtest.FirstBatch, bson.D{
			{"ok", 1},
			{"country", manufacturer.Country},
			{"mfrCommonName", manufacturer.MfrCommonName},
			{"mfrID", manufacturer.MfrID},
			{"mfrName", manufacturer.MfrName},
			{"vehicleTypes", bson.A{
				bson.D{
					{"isPrimary", manufacturer.VehicleTypes[0].IsPrimary},
					{"name", manufacturer.VehicleTypes[0].Name},
				},
			},
			}},
		))

		err := manufacturerRepo.UpdateManufacturer(context.Background(), manufacturerCollection, manufacturer.MfrID, manufacturer)
		assert.Nil(t, err)

	})

	mt.Run("failure", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		mt.AddMockResponses(bson.D{{"ok", 0}})

		err := manufacturerRepo.UpdateManufacturer(context.Background(), manufacturerCollection, manufacturer.MfrID, manufacturer)
		assert.Error(t, err)
	})

}

func TestGetAllManufacturers(t *testing.T) {
	manufacturerRepo := NewManufacturerRepo()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	manufacturer := models.Manufacturer{
		Country:       "UNITED STATES (USA)",
		MfrCommonName: "Tesla",
		MfrID:         955,
		MfrName:       "TESLA, INC.",
		VehicleTypes: []models.VehicleType{
			{
				IsPrimary: true,
				Name:      "Passenger Car",
			},
		},
	}

	mt.Run("success", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		first := mtest.CreateCursorResponse(1, "vehiclesapp.vehicles", mtest.FirstBatch, bson.D{
			{"ok", 1},
			{"country", manufacturer.Country},
			{"mfrCommonName", manufacturer.MfrCommonName},
			{"mfrID", manufacturer.MfrID},
			{"mfrName", manufacturer.MfrName},
			{"vehicleTypes", bson.A{
				bson.D{
					{"isPrimary", manufacturer.VehicleTypes[0].IsPrimary},
					{"name", manufacturer.VehicleTypes[0].Name},
				},
			},
			}},
		)
		killCursors := mtest.CreateCursorResponse(0, "vehiclesapp.vehicles", mtest.NextBatch)
		mt.AddMockResponses(first, killCursors)

		manufacturerResponse, err := manufacturerRepo.GetAllManufacturers(context.Background(), manufacturerCollection)
		assert.Nil(t, err)

		assert.Equal(t, []models.Manufacturer{
			manufacturer,
		}, manufacturerResponse)

	})

	mt.Run("failure", func(mt *mtest.T) {
		manufacturerCollection := mt.Coll

		mt.AddMockResponses(bson.D{{"ok", 0}})

		_, err := manufacturerRepo.GetManufacturer(context.Background(), manufacturerCollection, manufacturer.MfrID)
		assert.Error(t, err)
	})

}
