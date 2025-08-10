package rest

import (
	"errors"
	"net/http"
	"strconv"
	"sync"
	"userservice/internal/userService/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var lock = sync.Mutex{}

type CrudHandler[T any] struct {
	service *service.CrudService[T]
}

func NewCrudHandler[T any](s *service.CrudService[T]) *CrudHandler[T] {
	return &CrudHandler[T]{service: s}
}

func (h *CrudHandler[T]) CreateEntity(c echo.Context) error {
	// Evito problemas de concurrencia
	lock.Lock()
	defer lock.Unlock()

	var entity T

	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	createdEntity, err := h.service.RegisterEntity(&entity)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, createdEntity)
}

func (h *CrudHandler[T]) GetAllEntities(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	entities, err := h.service.FetchEntities()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, entities)
}

func (h *CrudHandler[T]) GetEntityByID(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	entity, err := h.service.FetchEntity(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, entity)
}

func (h *CrudHandler[T]) UpdateEntity(c echo.Context) error {
	// Evito problemas de concurrencia
	lock.Lock()
	defer lock.Unlock()

	var entity T

	if err := c.Bind(&entity); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "invalid request body",
		})
	}

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	updatedEntity, err := h.service.PatchEntity(&entity, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusAccepted, updatedEntity)
}

func (h *CrudHandler[T]) DeleteEntity(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)

	deleted, err := h.service.DeleteEntity(id)

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": err.Error(),
		})
	}

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusNoContent, deleted)
}
