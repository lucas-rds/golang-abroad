package api

// NewAgencyRouter returns a router with all agency endpoints
func (api API) routes() {
	api.GET("/agency", filterAgency(api.AgencyRepository))
	api.GET("/agency/:id", getAgency(api.AgencyRepository))
	api.POST("/agency", postAgency(api.AgencyRepository))
}
