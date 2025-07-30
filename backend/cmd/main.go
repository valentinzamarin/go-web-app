package main

import (
	"backend/internal/app/services"
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/postgres"
	postgresuser "backend/internal/infrastructure/postgres/postgres_user"
	"backend/internal/interfaces/api/rest"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	cfg := config.NewConfig()

	db, _ := postgres.NewConnection(cfg.DatabaseURL)

	if err := postgres.Migrate(db); err != nil {
		log.Fatalf("%v", err)
	}

	userRepo := postgresuser.NewGormUserRepository(db)

	userService := services.NewUserService(userRepo)

	e := echo.New()

	rest.NewUserController(e, userService)

	e.Logger.Fatal(e.Start(":8080"))
}
