package usecases

import (
	"github.com/go-foward/abroad/domain/agency/entities"
	"github.com/go-foward/abroad/domain/agency/repository"
)

type AgencyUseCases interface {
	CreateAgency(agency *entities.Agency) (*entities.Agency, error)
	GetAgencyById(id int) (*entities.Agency, error)
}

type AgencyUseCase struct {
	agencyRepository repository.AgencyRepository
}

func NewAgencyUseCase(repo repository.AgencyRepository) AgencyUseCases {
	return &AgencyUseCase{
		agencyRepository: repo,
	}
}

func (usecase AgencyUseCase) CreateAgency(agency *entities.Agency) (*entities.Agency, error) {
	return usecase.agencyRepository.Create(*agency)
}

func (usecase AgencyUseCase) GetAgencyById(id int) (*entities.Agency, error) {
	return usecase.agencyRepository.Get(id)
}
