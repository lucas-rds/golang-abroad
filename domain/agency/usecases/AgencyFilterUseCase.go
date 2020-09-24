package usecases

import (
	"github.com/go-foward/abroad/domain/agency/repository"
)

// AgencyFilterRequest ...
type AgencyFilterRequest struct {
	Filters map[string][]string
}

// AgencyFiltered ...
type AgencyFiltered struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// AgencyFilterResponse ...
type AgencyFilterResponse struct {
	Agencies []AgencyFiltered `json:"agencies"`
}

// AgencyFilterUseCaseExecuter ...
type AgencyFilterUseCaseExecuter interface {
	Execute(request AgencyGetByIDRequest) (*AgencyFilterResponse, error)
}

// AgencyFilterUseCase ...
type AgencyFilterUseCase struct {
	agencyRepository repository.AgencyRepository
}

// NewAgencyFilterUseCase ...
func NewAgencyFilterUseCase(repo repository.AgencyRepository) *AgencyFilterUseCase {
	return &AgencyFilterUseCase{
		agencyRepository: repo,
	}
}

// Execute ...
func (usecase AgencyFilterUseCase) Execute(request AgencyFilterRequest) (*AgencyFilterResponse, error) {

	dbAgencies, err := usecase.agencyRepository.Filter(request.Filters)
	if err != nil {
		return nil, err
	}

	agencies := []AgencyFiltered{}
	for _, value := range dbAgencies {
		agencies = append(agencies, AgencyFiltered{
			ID:   value.ID,
			Name: value.Name,
		})
	}
	response := &AgencyFilterResponse{
		Agencies: agencies,
	}
	return response, nil
}
