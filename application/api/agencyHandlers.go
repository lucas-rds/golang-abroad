package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-foward/abroad/application/agency/models"
	"github.com/labstack/echo"
)

// NewAgencyRouter returns a router with all agency endpoints
func (api API) enableAgencyHandlers() {
	api.GET("/agency", api.filterAgency)
	api.GET("/agency/:id", api.getAgency)
	api.POST("/agency", api.postAgency)
}

func (api API) filterAgency(c echo.Context) error {
	return c.String(http.StatusOK, "")
}

func (api API) getAgency(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	agency, err := api.AgencyController.GetAgency(id)
	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, agency)
}

func (api API) postAgency(c echo.Context) error {
	agencyRequestModel := new(models.AgencyRequest)

	err := c.Bind(agencyRequestModel)
	if err != nil {
		return err
	}

	createdAgency, err := api.AgencyController.CreateNewAgency(agencyRequestModel)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createdAgency)
}
