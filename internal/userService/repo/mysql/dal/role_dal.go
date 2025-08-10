package dal

import (
	"userservice/internal/userService/repo/mysql/dao"
	"userservice/pkg/models"

	"gorm.io/gorm"
)

type roleDAL struct {
	*crudDAL[models.Role]
}

func NewRoleDal(db *gorm.DB) dao.RoleDAO {
	return &roleDAL{
		crudDAL: &crudDAL[models.Role]{DB: db},
	}
}
