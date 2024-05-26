package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
)

type LocationMysql struct {
	db *sqlx.DB
}

func NewLocationMysql(db *sqlx.DB) *LocationMysql {
	return &LocationMysql{db: db}
}

func (r *LocationMysql) Create(location AlcoSafe.Location) (int, error) {
	query := "INSERT INTO Location (Country, City, Address, PostCode) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, location.Country, location.City, location.Address, location.PostCode)
	if err != nil {
		return 0, err
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	id := int(id64)
	return id, nil
}

func (r *LocationMysql) GetByID(locationID int) (AlcoSafe.Location, error) {
	var location AlcoSafe.Location
	query := "SELECT LocationID, Country, City, Address, PostCode FROM Location WHERE LocationID = ?"
	err := r.db.Get(&location, query, locationID)
	if err != nil {
		return location, err
	}
	return location, nil
}

func (r *LocationMysql) Delete(locationID int) error {
	query := "DELETE FROM Location WHERE LocationID = ?"
	_, err := r.db.Exec(query, locationID)
	return err
}
