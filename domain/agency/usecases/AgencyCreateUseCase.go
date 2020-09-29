package usecases

import (
	"github.com/go-foward/abroad/domain/agency/dbmodels"
	"github.com/go-foward/abroad/domain/agency/repository"
)

// AgencyCreateRequest ...
type AgencyCreateRequest struct {
	Name string
}

// AgencyCreated ...
type AgencyCreated struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// AgencyCreateResponse ...
type AgencyCreateResponse struct {
	Agency AgencyCreated `json:"agency"`
}

// AgencyCreateUseCase ...
type AgencyCreateUseCase struct {
	agencyRepository repository.AgencyRepository
}

// NewAgencyCreateUseCase ...
func NewAgencyCreateUseCase(repo repository.AgencyRepository) *AgencyCreateUseCase {
	return &AgencyCreateUseCase{
		agencyRepository: repo,
	}
}

// Execute ...
func (usecase AgencyCreateUseCase) Execute(request AgencyCreateRequest) (*AgencyCreateResponse, error) {
	dbAgency := dbmodels.NewAgencyDatabaseModel(request.Name)
	agency, err := usecase.agencyRepository.Create(*dbAgency)
	if err != nil {
		return nil, err
	}

	response := &AgencyCreateResponse{
		Agency: AgencyCreated{
			ID:   agency.ID,
			Name: agency.Name,
		},
	}
	return response, nil
}
