package usecases

import (
	"github.com/go-foward/abroad/domain/agency/repository"
)

type GetAgencyByIdRequest struct {
	id int
}

type GetAgencyResponse struct {
	Agency struct {
		Name string
	}
}

func NewGetAgencyByIdRequest(id int) *GetAgencyByIdRequest {
	return &GetAgencyByIdRequest{
		id: id,
	}
}

type GetAgencyByIdUseCaseExecuter interface {
	Execute(request GetAgencyByIdRequest) (*GetAgencyResponse, error)
}

type GetAgencyByIdUseCase struct {
	agencyRepository repository.AgencyRepository
}

func NewGetAgencyByIdUseCase(repo repository.AgencyRepository) *GetAgencyByIdUseCase {
	return &GetAgencyByIdUseCase{
		agencyRepository: repo,
	}
}

func (usecase GetAgencyByIdUseCase) Execute(request GetAgencyByIdRequest) (*GetAgencyResponse, error) {
	var response *GetAgencyResponse

	agency, err := usecase.agencyRepository.Get(request.id)
	if err != nil {
		return nil, err
	}

	response = &GetAgencyResponse{
		Agency: struct{ Name string }{Name: agency.Name},
	}
	return response, nil
}
