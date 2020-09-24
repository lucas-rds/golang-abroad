package api

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-foward/abroad/domain/agency/repository"
	"github.com/go-foward/abroad/domain/agency/usecases"
	"github.com/labstack/echo"
)

type Teste struct {
	Name string `json:"name"`
}

func (api API) filterAgency(repo repository.AgencyRepository) echo.HandlerFunc {
	return func(c echo.Context) error {

		request := usecases.AgencyFilterRequest{
			Filters: c.Request().URL.Query(),
		}

		log.Println(request)

		response, err := usecases.NewAgencyFilterUseCase(repo).Execute(request)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, response.Agencies)
	}
}

func (api API) getAgency(repo repository.AgencyRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))
		request := usecases.AgencyGetByIDRequest{ID: id}

		agency, err := usecases.NewAgencyGetByIDUseCase(repo).Execute(request)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, agency)
	}
}

func (api API) postAgency(repo repository.AgencyRepository) echo.HandlerFunc {
	return func(c echo.Context) error {
		var createAgencyRequest usecases.AgencyCreateRequest

		err := c.Bind(&createAgencyRequest)
		if err != nil {
			return err
		}

		usecase := usecases.NewAgencyCreateUseCase(repo)
		createdAgency, err := usecase.Execute(createAgencyRequest)

		if err != nil {
			return err
		}

		return c.JSON(http.StatusCreated, createdAgency)
	}
}
