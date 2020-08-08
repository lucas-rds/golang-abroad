package usecases

import (
	"github.com/go-foward/abroad/domain/agency/entities"
	"github.com/go-foward/abroad/domain/agency/repository"
)

type AgencyUseCases interface {
	CreateAgency(agency entities.Agency)
	GetAgencyById(id int) *entities.Agency
}

type AgencyUseCase struct {
	agencyRepository repository.AgencyRepository
}

func NewAgencyUseCase(repo repository.AgencyRepository) AgencyUseCases {
	return &AgencyUseCase{
		agencyRepository: repo,
	}
}

func (usecase AgencyUseCase) CreateAgency(agency entities.Agency) {
	usecase.agencyRepository.Save(agency)
}

func (usecase AgencyUseCase) GetAgencyById(id int) *entities.Agency {
	return usecase.agencyRepository.Get(id)
}
