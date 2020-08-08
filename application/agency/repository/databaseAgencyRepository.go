package repository

import (
	"log"

	"github.com/go-foward/abroad/domain/agency/entities"
	"github.com/jmoiron/sqlx"
)

type DatabaseAgencyRepository struct {
	Database *sqlx.DB
}

func NewDatabaseAgencyRepository(db *sqlx.DB) *DatabaseAgencyRepository {
	return &DatabaseAgencyRepository{
		Database: db,
	}
}

func (repo DatabaseAgencyRepository) Get(id int) *entities.Agency {
	var agency entities.Agency
	err := repo.Database.Get(&agency, "SELECT * FROM agency WHERE id = $1", id)
	if err != nil {
		log.Fatal(err)
	}
	return &agency
}

func (repo DatabaseAgencyRepository) Save(agency entities.Agency) {}

func (repo DatabaseAgencyRepository) Create(agency entities.Agency) {}
