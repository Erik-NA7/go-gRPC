package controllers

import (
	"context"
	"database/sql"
	"log"
	"unsia/models"
	"unsia/pb/cities"
)

// City Struct
type City struct{
	DB *sql.DB
	Log *log.Logger
	cities.UnimplementedCitiesServiceServer
}

// GetCity function
func (s *City) GetCity(ctx context.Context, in *cities.Id) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Get(ctx, s.DB, in)
	return &cityModel.Pb, err
}

// CreateCity function
func (s *City) Create(ctx context.Context, in *cities.CityInput) (*cities.City, error) {
	var cityModel models.City
	err := cityModel.Create(ctx, s.DB, in)
	return &cityModel.Pb, err
}