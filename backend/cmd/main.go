package main

import (
	"backend/internal/app/services"
	"backend/internal/infrastructure/config"
	"backend/internal/infrastructure/postgres"
	postgresuser "backend/internal/infrastructure/postgres/postgres_user"
	"backend/internal/interfaces/api/rest"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	/*
		ломаются корсы на Vue части
		пока так, после разберусь, что и как
	*/

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173", "http://localhost:8081"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	rest.NewUserController(e, userService)

	e.Logger.Fatal(e.Start(":8080"))
}
