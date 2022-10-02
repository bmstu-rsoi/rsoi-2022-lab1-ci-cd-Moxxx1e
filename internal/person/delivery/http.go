package delivery

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"

	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/person/repository"
)

const (
	apiPrefix = "/api/v1"

	locationValueFormat = "/api/v1/persons/%d"
)

type Handler struct {
	usecase usecase
}

func NewHandler(u usecase) *Handler {
	return &Handler{usecase: u}
}

func (h *Handler) Configure(e *echo.Echo) {
	e.POST(apiPrefix+"/persons", h.Create())
	e.DELETE(apiPrefix+"/persons/:id", h.Delete())
	e.PATCH(apiPrefix+"/persons/:id", h.Update())
	e.GET(apiPrefix+"/persons/:id", h.GetByID())
	e.GET(apiPrefix+"/persons", h.GetAll())
}

func (h *Handler) Create() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := &person{}

		if err := ctx.Bind(req); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}

		id, err := h.usecase.Create(context.Background(), req.toModel())
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		locationValue := fmt.Sprintf(locationValueFormat, id)
		ctx.Response().Header().Set("Location", locationValue)

		return ctx.JSON(http.StatusCreated, nil)
	}
}

func (h *Handler) Delete() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		strID := ctx.Param("id")
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, nil)
		}

		err = h.usecase.Delete(context.Background(), id)
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusNoContent, nil)
	}
}

func (h *Handler) GetByID() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, nil)
		}

		model, err := h.usecase.GetByID(context.Background(), id)
		if err != nil {
			if errors.Is(err, repository.ErrNoPersonWithSuchID) {
				return ctx.JSON(http.StatusNotFound, httpError{Message: ""})
			}
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, fromModel(model))
	}
}

func (h *Handler) GetAll() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		persons, err := h.usecase.GetAll(context.Background())
		if err != nil {
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		var response []person
		for _, p := range *persons {
			response = append(response, fromModel(p))
		}

		return ctx.JSON(http.StatusOK, response)
	}
}

func (h *Handler) Update() echo.HandlerFunc {
	return func(ctx echo.Context) error {
		req := &person{}
		if err := ctx.Bind(req); err != nil {
			return ctx.JSON(http.StatusBadRequest, err)
		}
		id, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
		if err != nil {
			return ctx.JSON(http.StatusBadRequest, nil)
		}
		req.ID = id

		updated, err := h.usecase.Update(context.Background(), req.toModel())
		if err != nil {
			if errors.Is(err, repository.ErrNoPersonWithSuchID) {
				return ctx.JSON(http.StatusNotFound, httpError{Message: ""})
			}
			return ctx.JSON(http.StatusInternalServerError, err)
		}

		return ctx.JSON(http.StatusOK, fromModel(updated))
	}
}
