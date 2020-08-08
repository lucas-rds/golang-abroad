package main

import (
	"log"

	"github.com/go-foward/abroad/application/agency/controller"
	"github.com/go-foward/abroad/application/agency/repository"
	"github.com/go-foward/abroad/application/api"
	"github.com/go-foward/abroad/domain/agency/usecases"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
)

func main() {
	db, err := sqlx.Connect("pgx", "postgres://postgres:adm123@localhost:5432/abroad?sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}

	agencyRepository := repository.NewAgencyDatabaseRepository(db)
	agencyUseCases := usecases.NewAgencyUseCase(agencyRepository)
	agencyController := controller.NewAgencyController(agencyRepository, agencyUseCases)

	v1 := api.NewAPI(agencyController)
	v1.Run()
}
