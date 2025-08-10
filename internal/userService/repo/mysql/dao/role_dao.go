package dao

import "userservice/pkg/models"

type RoleDAO interface {
	CrudDAO[models.Role]
}
