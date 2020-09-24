package repository

import (
	"fmt"
	"log"
	"strings"

	"github.com/go-foward/abroad/domain/agency/dbmodels"
	"github.com/jmoiron/sqlx"
)

// AgencyDatabaseRepository ...
type AgencyDatabaseRepository struct {
	Database *sqlx.DB
}

// NewAgencyDatabaseRepository ...
func NewAgencyDatabaseRepository(db *sqlx.DB) *AgencyDatabaseRepository {
	return &AgencyDatabaseRepository{
		Database: db,
	}
}

// Get ...
func (repo AgencyDatabaseRepository) Get(id int) (*dbmodels.AgencyDatabaseModel, error) {
	var agency dbmodels.AgencyDatabaseModel
	err := repo.Database.Get(&agency, "SELECT * FROM agency WHERE id = $1", id)
	return &agency, err
}

// Filter ...
func (repo AgencyDatabaseRepository) Filter(filters map[string][]string) ([]dbmodels.AgencyDatabaseModel, error) {
	var agencies []dbmodels.AgencyDatabaseModel

	where := "WHERE 1=1"
	for key, value := range filters {
		where = fmt.Sprintf("%s AND %s = %s", where, key, fmt.Sprintf("'%s'", strings.Join(value, "','")))
	}
	log.Println(fmt.Sprintf("SELECT * FROM agency %s", where))
	err := repo.Database.Select(&agencies, fmt.Sprintf("SELECT * FROM agency %s", where))
	log.Println(agencies)

	return agencies, err
}

// Save ...
func (repo AgencyDatabaseRepository) Save(agency dbmodels.AgencyDatabaseModel) (*dbmodels.AgencyDatabaseModel, error) {
	return nil, nil
}

// Create ...
func (repo AgencyDatabaseRepository) Create(agency dbmodels.AgencyDatabaseModel) (*dbmodels.AgencyDatabaseModel, error) {
	preparedStatement, err := repo.Database.PrepareNamed("INSERT INTO agency(name) VALUES(:name) RETURNING id")
	if err != nil {
		return nil, err
	}

	var id int
	err = preparedStatement.Get(&id, agency)
	if err != nil {
		return nil, err
	}

	createdAgency, err := repo.Get(id)
	if err != nil {
		return nil, err
	}

	return createdAgency, err
}
