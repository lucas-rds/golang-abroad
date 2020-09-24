package dbmodels

// AgencyDatabaseModel model
type AgencyDatabaseModel struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
}

// AgencyDatabaseModelOption type
type AgencyDatabaseModelOption func(*AgencyDatabaseModel)

// NewAgencyDatabaseModel AgencyDatabaseModel builder
func NewAgencyDatabaseModel(name string, options ...AgencyDatabaseModelOption) *AgencyDatabaseModel {
	model := &AgencyDatabaseModel{
		Name: name,
	}
	for _, option := range options {
		option(model)
	}
	return model
}

// AgencyID adds id into model
func AgencyID(id int) AgencyDatabaseModelOption {
	return func(model *AgencyDatabaseModel) {
		model.ID = id
	}
}
