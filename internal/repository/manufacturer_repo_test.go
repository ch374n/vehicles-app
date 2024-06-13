package repository

import (
	"context"
	"errors"
	"testing"
	models "github.com/ch374n/vehicles-app/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ManufacturerSuite struct {
	suite.Suite
	*require.Assertions
	ctrl              *gomock.Controller
	collections *mongo.Collection
	manufacturerRepo ManufacturerRepo
}

func TestManufacturerSuite(t *testing.T) {
	suite.Run(t, new(ManufacturerSuite))
}

func (s *ManufacturerSuite) SetupTest() {
	s.Assertions = require.New(s.T())

	s.ctrl = gomock.NewController(s.T())

	s.collections = //TODO

	s.manufacturerRepo = NewManufacturerRepo()
}

func (s *ManufacturerSuite) TearDownTest() {
	s.ctrl.Finish()
}

func (s *ManufacturerSuite) TestCreateManufacturer() {
	ctx := context.Background()

	manufacturerInput := models.Manufacturer{
		Country:        "UNITED STATES (USA)",
		MfrCommonName:  "Tesla",
		MfrID:          955,
		MfrName:        "TESLA, INC.",
		VehicleTypes: []models.VehicleType{
			{
				IsPrimary: true,
				Name:      "Passenger Car",
			},
		},
	}

	s.collections.EXPECT().Find(ctx, gomock.Eq(manufacturerInput)).Return(nil)

	err := s.manufacturerRepo.CreateManufacturer(ctx, manufacturerInput)

	s.NoError(err)


	// test error scenario
	s.collections.EXPECT().Find(ctx, gomock.Eq(manufacturerInput)).Return(errors.New("database error"))
	err = s.manufacturerRepo.CreateManufacturer(ctx, manufacturerInput)
	s.Error(err)

}


func (s *ManufacturerSuite) TestGetAllManufacturers() {
	ctx := context.Background()

	expectedManufacturers := []models.Manufacturer{
		{
			Country:        "UNITED STATES (USA)",
			MfrCommonName:  "Tesla",
			MfrID:          955,
			MfrName:        "TESLA, INC.",
			VehicleTypes: []models.VehicleType{
				{
					IsPrimary: true,
					Name:      "Passenger Car",
				},
			},
		},
		{
			Country:        "UNITED STATES (USA)",
			MfrCommonName:  "Ford",
			MfrID:          474,
			MfrName:        "FORD MOTOR COMPANY",
			VehicleTypes: []models.VehicleType{
				{
					IsPrimary: true,
					Name:      "Passenger Car",
				},
				{
					IsPrimary: false,
					Name:      "Truck",
				},
			},
		},
	}

	s.manufacturerRepo.EXPECT().GetAllManufacturers(ctx).Return(expectedManufacturers, nil)

	manufacturers, err := s.manufacturerRepo.GetAllManufacturers(ctx)

	s.NoError(err)
	s.Equal(expectedManufacturers, manufacturers)

	// Test error scenario
	s.manufacturerRepo.EXPECT().GetAllManufacturers(ctx).Return(nil, errors.New("database error"))
	_, err = s.manufacturerRepo.GetAllManufacturers(ctx)
	s.Error(err)
}

func (s *ManufacturerSuite) TestUpdateManufacturer() {
	ctx := context.Background()
	id := 955

	manufacturerInput := models.Manufacturer{
		Country:        "UNITED STATES (USA)",
		MfrCommonName:  "Tesla",
		MfrID:          955,
		MfrName:        "TESLA, INC.",
		VehicleTypes: []models.VehicleType{
			{
				IsPrimary: true,
				Name:      "Passenger Car",
			},
		},
	}

	s.manufacturerRepo.EXPECT().UpdateManufacturer(ctx, id, gomock.Eq(manufacturerInput)).Return(nil)

	err := s.manufacturerRepo.UpdateManufacturer(ctx, id, manufacturerInput)

	s.NoError(err)

	// Test error scenario
	s.manufacturerRepo.EXPECT().UpdateManufacturer(ctx, id, gomock.Any()).Return(errors.New("database error"))
	err = s.manufacturerRepo.UpdateManufacturer(ctx, id, manufacturerInput)
	s.Error(err)
}

func (s *ManufacturerSuite) TestDeleteManufacturer() {
	ctx := context.Background()
	id := 955

	s.manufacturerRepo.EXPECT().DeleteManufacturer(ctx, id).Return(nil)

	err := s.manufacturerRepo.DeleteManufacturer(ctx, id)

	s.NoError(err)

	// Test error scenario
	s.manufacturerRepo.EXPECT().DeleteManufacturer(ctx, id).Return(errors.New("database error"))
	err = s.manufacturerRepo.DeleteManufacturer(ctx, id)
	s.Error(err)
}