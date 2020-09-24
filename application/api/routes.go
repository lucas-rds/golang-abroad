package api

// NewAgencyRouter returns a router with all agency endpoints
func (api API) routes() {
	api.GET("/agency", api.filterAgency(api.AgencyRepository))
	api.GET("/agency/:id", api.getAgency(api.AgencyRepository))
	api.POST("/agency", api.postAgency(api.AgencyRepository))
}
