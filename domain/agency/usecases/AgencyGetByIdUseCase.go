package usecases

import (
	"github.com/go-foward/abroad/domain/agency/repository"
)

// AgencyGetByIDRequest ...
type AgencyGetByIDRequest struct {
	ID int
}

// AgencyRetrieved ...
type AgencyRetrieved struct {
	Name string `json:"name"`
}

// AgencyGetResponse ...
type AgencyGetResponse struct {
	Agency AgencyRetrieved `json:"agency"`
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
		Agency: AgencyRetrieved{Name: agency.Name},
	}
	return response, nil
}
