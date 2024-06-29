package models

import (
	"context"
	"database/sql"
	"unsia/pb/cities"
	"fmt"
	"log"
)

type City struct {
	Pb cities.City
	Log *log.Logger
}

func (u *City) Get(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `SELECT * FROM cities WHERE id=$1;`
	err := db.QueryRowContext(ctx, query, in.Id).Scan(&u.Pb.Id, &u.Pb.Name) 
	if err != nil {
		u.Log.Println("Error on query: ", err)
		return err
	}
	return nil
}

func (u *City) Create(ctx context.Context, db *sql.DB, in *cities.CityInput) error {
	query := `INSERT INTO cities (name) VALUES ($1) RETURNING id;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		u.Log.Println("Error on query: ", err)
		return err
	}
	
	err = stmt.QueryRowContext(ctx, in.Name).Scan(&u.Pb.Id)
	
	if err != nil {
		u.Log.Println("Error on query: ", err)
		return err
	}

	return nil
}

func (u *City) Update(ctx context.Context, db *sql.DB, in *cities.City) error {
	query := `UPDATE cities SET name=$2 WHERE id=$1 RETURNING id, name;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		u.Log.Println("Error on query: ", err)
		return err
	}
	
	err = stmt.QueryRowContext(ctx, int32(in.Id), in.Name).Scan(&u.Pb.Id, &u.Pb.Name)
	
	if err != nil {
		u.Log.Println("Error on query: ", err)
		return err
	}

	return nil
}

func (u *City) Delete(ctx context.Context, db *sql.DB, in *cities.Id) error {
	query := `DELETE FROM cities WHERE id=$1;`
	stmt, err := db.PrepareContext(ctx, query)
	if err != nil {
		u.Log.Println("Error on query: ", err)
		return err
	}

	rs, err := stmt.ExecContext(ctx, in.Id)
	
	if err != nil {
		u.Log.Println("Error on query: ", err)
		return err
	}
	
	affected, err := rs.RowsAffected()
	if err != nil {
		u.Log.Println("Query Result error: ", err)
		return err
	}

	if affected == 0 {
		u.Log.Printf("Query Result: Data with id = %d not found", in.Id)
		return fmt.Errorf("data not found")
	}
	
	return nil
}