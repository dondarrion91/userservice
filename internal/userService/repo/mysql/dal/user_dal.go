package dal

import (
	"userservice/internal/userService/repo/mysql/dao"
	"userservice/pkg/models"

	"gorm.io/gorm"
)

type userDAL struct {
	*crudDAL[models.User]
}

func NewUserDAL(db *gorm.DB) dao.UserDAO {
	return &userDAL{
		crudDAL: &crudDAL[models.User]{DB: db},
	}
}
