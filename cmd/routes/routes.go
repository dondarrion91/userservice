package routes

import (
	"net/http"
	"userservice/internal/userService/repo/mysql/dal"
	"userservice/internal/userService/repo/mysql/dao"
	"userservice/internal/userService/rest"
	"userservice/internal/userService/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func crud[T any](group *echo.Group, dal dao.CrudDAO[T]) {
	crudService := service.NewCrudService(dal)
	crudHandler := rest.NewCrudHandler(crudService)

	group.POST("", crudHandler.CreateEntity)
	group.GET("", crudHandler.GetAllEntities)
	group.GET("/:id", crudHandler.GetEntityByID)
	group.PATCH("/:id", crudHandler.UpdateEntity)
	group.DELETE("/:id", crudHandler.DeleteEntity)
}

func user(r *echo.Echo, db *gorm.DB) {
	userGroup := r.Group("/user")

	userDAL := dal.NewUserDAL(db)

	crud(userGroup, userDAL)
}

func role(r *echo.Echo, db *gorm.DB) {
	roleGroup := r.Group("/role")

	roleDal := dal.NewRoleDal(db)

	crud(roleGroup, roleDal)
}

func address(r *echo.Echo, db *gorm.DB) {
	addressGroup := r.Group("/address")

	addressDAL := dal.NewAddressDAL(db)

	crud(addressGroup, addressDAL)
}

func Routes(r *echo.Echo, db *gorm.DB) *echo.Echo {
	r.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "UP")
	})

	// Grupo de routes
	user(r, db)
	address(r, db)
	role(r, db)

	return r
}
