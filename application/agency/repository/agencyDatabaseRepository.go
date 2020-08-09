package repository

import (
	"github.com/go-foward/abroad/application/agency/models"
	"github.com/go-foward/abroad/domain/agency/entities"
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

func (repo AgencyDatabaseRepository) Get(id int) (*entities.Agency, error) {
	var agency entities.Agency
	err := repo.Database.Get(&agency, "SELECT * FROM agency WHERE id = $1", id)
	return &agency, err
}

func (repo AgencyDatabaseRepository) Save(agency entities.Agency) (*entities.Agency, error) {
	return nil, nil
}

func (repo AgencyDatabaseRepository) Create(agency entities.Agency) (*entities.Agency, error) {
	var id int
	dbAgency := models.DbAgencyFromDomain(&agency)
	preparedStatement, err := repo.Database.PrepareNamed("INSERT INTO agency(name) VALUES(:name) RETURNING id")
	if err != nil {
		return nil, err
	}

	err = preparedStatement.Get(&id, dbAgency)
	if err != nil {
		return nil, err
	}

	createdAgency, err := repo.Get(id)
	if err != nil {
		return nil, err
	}

	return createdAgency, err
}
