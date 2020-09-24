package usecases

import (
	"github.com/go-foward/abroad/domain/agency/repository"
)

// AgencyGetByIDRequest ...
type AgencyGetByIDRequest struct {
	ID int
}

// AgencyGet ...
type AgencyGet struct {
	Name string `json:"name"`
}

// AgencyGetResponse ...
type AgencyGetResponse struct {
	Agency AgencyGet `json:"agency"`
}

// AgencyGetByIDUseCaseExecuter ...
type AgencyGetByIDUseCaseExecuter interface {
	Execute(request AgencyGetByIDRequest) (*AgencyGetResponse, error)
}

// AgencyGetByIDUseCase ...
type AgencyGetByIDUseCase struct {
	agencyRepository repository.AgencyRepository
}

// NewAgencyGetByIDUseCase ...
func NewAgencyGetByIDUseCase(repo repository.AgencyRepository) *AgencyGetByIDUseCase {
	return &AgencyGetByIDUseCase{
		agencyRepository: repo,
	}
}

// Execute ...
func (usecase AgencyGetByIDUseCase) Execute(request AgencyGetByIDRequest) (*AgencyGetResponse, error) {
	var response *AgencyGetResponse

	agency, err := usecase.agencyRepository.Get(request.ID)
	if err != nil {
		return nil, err
	}

	response = &AgencyGetResponse{
		Agency: AgencyGet{Name: agency.Name},
	}
	return response, nil
}
