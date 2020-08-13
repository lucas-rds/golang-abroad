package dbmodels

// AgencyRequest model
type AgencyDatabaseModel struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}

type AgencyDatabaseModelOption func(*AgencyDatabaseModel)

//DbAgencyFromDomain ...
func NewAgencyDatabaseModel(name string, options ...AgencyDatabaseModelOption) *AgencyDatabaseModel {
	model := &AgencyDatabaseModel{
		Name: name,
	}
	for _, option := range options {
		option(model)
	}
	return model
}

func Id(id int) AgencyDatabaseModelOption {
	return func(model *AgencyDatabaseModel) {
		model.Id = id
	}
}
