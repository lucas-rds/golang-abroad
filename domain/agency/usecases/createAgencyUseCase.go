package usecases

import (
	"github.com/go-foward/abroad/domain/agency/dbmodels"
	"github.com/go-foward/abroad/domain/agency/repository"
)

type CreateAgencyRequest struct {
	Name string
}

type CreateAgencyResponse struct {
	Agency struct {
		Id   int
		Name string
	}
}

type CreateAgencyUseCaseExecuter interface {
	Execute(request CreateAgencyRequest) (*CreateAgencyResponse, error)
}

type CreateAgencyUseCase struct {
	agencyRepository repository.AgencyRepository
}

func NewCreateAgencyUseCase(repo repository.AgencyRepository) *CreateAgencyUseCase {
	return &CreateAgencyUseCase{
		agencyRepository: repo,
	}
}

func (usecase CreateAgencyUseCase) Execute(request CreateAgencyRequest) (*CreateAgencyResponse, error) {
	var response *CreateAgencyResponse

	dbAgency := dbmodels.NewAgencyDatabaseModel(request.Name)

	agency, err := usecase.agencyRepository.Create(*dbAgency)
	if err != nil {
		return nil, err
	}

	response = &CreateAgencyResponse{
		Agency: struct {
			Id   int
			Name string
		}{
			Id:   agency.Id,
			Name: agency.Name,
		},
	}
	return response, nil
}
