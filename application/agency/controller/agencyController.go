package controller

import (
	"github.com/go-foward/abroad/application/agency/models"
	"github.com/go-foward/abroad/domain/agency/entities"
	"github.com/go-foward/abroad/domain/agency/usecases"
	"github.com/go-foward/abroad/domain/repository"
)

type AgencyControl interface {
	CreateNewAgency(agency entities.Agency)
	GetAgency(id int) *models.AgencyResponse
}

type AgencyController struct {
	agencyRepository repository.AgencyRepository
	agencyUseCase    usecases.AgencyUseCases
}

func NewAgencyController(
	repository repository.AgencyRepository,
	agencyUseCase usecases.AgencyUseCases) AgencyControl {

	return &AgencyController{
		agencyRepository: repository,
		agencyUseCase:    agencyUseCase,
	}
}

func (controller AgencyController) CreateNewAgency(agency entities.Agency) {
	controller.agencyUseCase.CreateAgency(agency)
}

func (controller AgencyController) GetAgency(id int) *models.AgencyResponse {
	agency := controller.agencyUseCase.GetAgencyById(id)
	return models.AgencyResponseFromDomain(agency)
}
