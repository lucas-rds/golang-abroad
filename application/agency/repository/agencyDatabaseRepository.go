package repository

import (
	"github.com/go-foward/abroad/domain/agency/dbmodels"
	"github.com/jmoiron/sqlx"
)

type AgencyDatabaseRepository struct {
	Database *sqlx.DB
}

func NewAgencyDatabaseRepository(db *sqlx.DB) *AgencyDatabaseRepository {
	return &AgencyDatabaseRepository{
		Database: db,
	}
}

func (repo AgencyDatabaseRepository) Get(id int) (*dbmodels.AgencyDatabaseModel, error) {
	var agency dbmodels.AgencyDatabaseModel
	err := repo.Database.Get(&agency, "SELECT * FROM agency WHERE id = $1", id)
	return &agency, err
}

func (repo AgencyDatabaseRepository) Save(agency dbmodels.AgencyDatabaseModel) (*dbmodels.AgencyDatabaseModel, error) {
	return nil, nil
}

func (repo AgencyDatabaseRepository) Create(agency dbmodels.AgencyDatabaseModel) (*dbmodels.AgencyDatabaseModel, error) {
	var id int
	preparedStatement, err := repo.Database.PrepareNamed("INSERT INTO agency(name) VALUES(:name) RETURNING id")
	if err != nil {
		return nil, err
	}

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
