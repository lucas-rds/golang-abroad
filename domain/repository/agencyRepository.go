package repository

import "github.com/go-foward/abroad/domain/agency/entities"

// AgencyCreator interface
type AgencyCreator interface {
	Create(agency entities.Agency)
}

// AgencyEditor interface
type AgencyEditor interface {
	Save(agency entities.Agency)
}

// AgencyRetriever interface
type AgencyRetriever interface {
	Get(id int) *entities.Agency
}

// AgencyRepository agency interfaces composition
type AgencyRepository interface {
	AgencyCreator
	AgencyEditor
	AgencyRetriever
}
