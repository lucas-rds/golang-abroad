package controller

import (
	"github.com/go-foward/abroad/application/agency/models"
	"github.com/go-foward/abroad/domain/agency/repository"
	"github.com/go-foward/abroad/domain/agency/usecases"
)

type AgencyControl interface {
	CreateNewAgency(*models.AgencyRequest) (*models.AgencyResponse, error)
	GetAgency(id int) (*models.AgencyResponse, error)
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

func (controller AgencyController) CreateNewAgency(agency *models.AgencyRequest) (*models.AgencyResponse, error) {
	domainAgency := models.AgencyRequestToDomain(agency)
	createdAgency, err := controller.agencyUseCase.CreateAgency(domainAgency)
	return models.AgencyResponseFromDomain(createdAgency), err
}

func (controller AgencyController) GetAgency(id int) (*models.AgencyResponse, error) {
	agency, err := controller.agencyUseCase.GetAgencyById(id)
	return models.AgencyResponseFromDomain(agency), err
}
