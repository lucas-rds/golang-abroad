package repository

import (
	"log"

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

func (repo AgencyDatabaseRepository) Get(id int) *entities.Agency {
	var agency entities.Agency
	err := repo.Database.Get(&agency, "SELECT * FROM agency WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	return &agency
}

func (repo AgencyDatabaseRepository) Save(agency entities.Agency) {}

func (repo AgencyDatabaseRepository) Create(agency entities.Agency) {}
