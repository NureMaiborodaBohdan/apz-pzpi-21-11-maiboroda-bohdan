package repository

import (
	AlcoSafe "apz-pzpi-21-11-maiboroda-bohdan-task2"
	"github.com/jmoiron/sqlx"
)

type CompanyMysql struct {
	db *sqlx.DB
}

func NewCompanyMysql(db *sqlx.DB) *CompanyMysql {
	return &CompanyMysql{db: db}
}

func (r *CompanyMysql) Create(company AlcoSafe.Company) (int, error) {
	query := "INSERT INTO Company (Name, Description, LegalLimit, LocationID) VALUES (?, ?, ?, ?)"
	result, err := r.db.Exec(query, company.Name, company.Description, company.LegalLimit, company.LocationID)
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

func (r *CompanyMysql) GetAll() ([]AlcoSafe.Company, error) {
	var companies []AlcoSafe.Company
	query := "SELECT * FROM Company"
	err := r.db.Select(&companies, query)
	return companies, err
}

func (r *CompanyMysql) GetByID(companyID int) (AlcoSafe.Company, error) {
	var company AlcoSafe.Company
	query := "SELECT * FROM Company WHERE CompanyID = ?"
	err := r.db.Get(&company, query, companyID)
	return company, err
}

func (r *CompanyMysql) Delete(companyID int) error {
	query := "DELETE FROM Company WHERE CompanyID = ?"
	_, err := r.db.Exec(query, companyID)
	return err
}
