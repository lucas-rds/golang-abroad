package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-foward/abroad/domain/agency/usecases"
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
	request := usecases.NewGetAgencyByIdRequest(id)

	usecase := usecases.NewGetAgencyByIdUseCase(api.AgencyRepository)
	agency, err := usecase.Execute(*request)

	if err != nil {
		log.Println(err)
	}
	return c.JSON(http.StatusOK, agency)
}

func (api API) postAgency(c echo.Context) error {
	var createAgencyRequest usecases.CreateAgencyRequest

	err := c.Bind(&createAgencyRequest)
	if err != nil {
		return err
	}

	usecase := usecases.NewCreateAgencyUseCase(api.AgencyRepository)
	createdAgency, err := usecase.Execute(createAgencyRequest)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, createdAgency)
}
