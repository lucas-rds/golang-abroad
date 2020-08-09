package models

import (
	"github.com/go-foward/abroad/domain/agency/entities"
)

// AgencyRequest model
type AgencyDatabaseModel struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

//ToDomain ...
func DbAgencyFromDomain(agency *entities.Agency) *AgencyDatabaseModel {
	return &AgencyDatabaseModel{
		Name: agency.Name,
	}
}
