package api

func (api API) routes() {
	api.GET("/agency", filterAgency(api.AgencyRepository))
	api.GET("/agency/:id", getAgency(api.AgencyRepository))
	api.POST("/agency", postAgency(api.AgencyRepository))
}
