package models

import (
	"github.com/go-foward/abroad/domain/agency/entities"
)

// AgencyRequest model
type AgencyDatabaseModel struct {
	Name string
}

//ToDomain ...
func DbAgencyFromDomain(agency *entities.Agency) *AgencyDatabaseModel {
	return &AgencyDatabaseModel{
		Name: agency.Name,
	}
}
