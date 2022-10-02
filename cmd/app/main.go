package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/person/delivery"
	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/person/repository"
	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/person/usecase"
)

const (
	dsn = "postgres://oskolganov:postgres@localhost:5432/postgres?sslmode=disable"
)

func main() {
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal(fmt.Errorf("error connecting to database: %w", err))
	}

	repo := repository.NewPG(db)
	uc := usecase.New(repo)
	handler := delivery.NewHandler(uc)

	e := echo.New()
	handler.Configure(e)

	log.Fatal(e.Start(":8890"))
}
