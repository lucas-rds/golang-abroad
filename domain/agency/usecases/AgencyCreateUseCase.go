package usecases

import (
	"github.com/go-foward/abroad/domain/agency/dbmodels"
	"github.com/go-foward/abroad/domain/agency/repository"
)

// AgencyCreateRequest ...
type AgencyCreateRequest struct {
	Name string
}

// AgencyCreateResponse ...
type AgencyCreateResponse struct {
	Agency struct {
		ID   int
		Name string
	}
}

// AgencyCreateUseCaseExecuter ...
type AgencyCreateUseCaseExecuter interface {
	Execute(request AgencyCreateRequest) (*AgencyCreateResponse, error)
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
		Agency: struct {
			ID   int
			Name string
		}{
			ID:   agency.ID,
			Name: agency.Name,
		},
	}
	return response, nil
}
