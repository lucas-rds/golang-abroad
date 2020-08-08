package api

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

// NewAgencyRouter returns a router with all agency endpoints
func (api API) enableAgencyHandlers() {
	api.GET("/agency", api.filterAgency)
	api.GET("/agency/:id", api.getAgency)
	api.POST("/agency", api.postAgency)
}

func (api API) filterAgency(c echo.Context) error {
	return c.String(http.StatusOK, c.Get("requestID").(string))

}

func (api API) getAgency(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	return c.JSON(http.StatusOK, api.AgencyController.GetAgency(id))
}

func (api API) postAgency(c echo.Context) error {
	return c.String(http.StatusOK, "POST Agency")
}
