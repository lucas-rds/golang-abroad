package models

import (
	"github.com/go-foward/abroad/domain/agency/entities"
)

// AgencyRequest model
type AgencyRequest struct {
	Name string
}

// AgencyResponse model
type AgencyResponse struct {
	Name string `json:"name"`
}

//ToDomain ...
func ToDomain(agencyRequest *AgencyRequest) *entities.Agency {
	return &entities.Agency{
		Name: agencyRequest.Name,
	}
}

//FromDomain ...
func FromDomain(agency *entities.Agency) *AgencyResponse {
	return &AgencyResponse{
		Name: agency.Name,
	}
}
