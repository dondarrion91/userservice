package dal

import (
	"userservice/internal/userService/repo/mysql/dao"
	"userservice/pkg/models"

	"gorm.io/gorm"
)

type addressDAL struct {
	*crudDAL[models.Address]
}

func NewAddressDAL(db *gorm.DB) dao.AddressDAO {
	return &addressDAL{
		crudDAL: &crudDAL[models.Address]{DB: db},
	}
}
