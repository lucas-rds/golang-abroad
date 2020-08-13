package repository

import "github.com/go-foward/abroad/domain/agency/dbmodels"

// AgencyCreator interface
type AgencyCreator interface {
	Create(agency dbmodels.AgencyDatabaseModel) (*dbmodels.AgencyDatabaseModel, error)
}

// AgencyEditor interface
type AgencyEditor interface {
	Save(agency dbmodels.AgencyDatabaseModel) (*dbmodels.AgencyDatabaseModel, error)
}

// AgencyRetriever interface
type AgencyRetriever interface {
	Get(id int) (*dbmodels.AgencyDatabaseModel, error)
}

// AgencyRepository agency interfaces composition
type AgencyRepository interface {
	AgencyCreator
	AgencyEditor
	AgencyRetriever
}
