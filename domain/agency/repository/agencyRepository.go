package repository

import "github.com/go-foward/abroad/domain/agency/entities"

// AgencyCreator interface
type AgencyCreator interface {
	Create(agency entities.Agency) (*entities.Agency, error)
}

// AgencyEditor interface
type AgencyEditor interface {
	Save(agency entities.Agency) (*entities.Agency, error)
}

// AgencyRetriever interface
type AgencyRetriever interface {
	Get(id int) (*entities.Agency, error)
}

// AgencyRepository agency interfaces composition
type AgencyRepository interface {
	AgencyCreator
	AgencyEditor
	AgencyRetriever
}
