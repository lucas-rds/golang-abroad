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

// AgencyRequestToDomain parses request to be used in domain
func AgencyRequestToDomain(agencyRequest *AgencyRequest) *entities.Agency {
	return &entities.Agency{
		Name: agencyRequest.Name,
	}
}

// AgencyResponseFromDomain parses domain into a request
func AgencyResponseFromDomain(agency *entities.Agency) *AgencyResponse {
	return &AgencyResponse{
		Name: agency.Name,
	}
}
