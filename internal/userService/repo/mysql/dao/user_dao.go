package dao

import "userservice/pkg/models"

type UserDAO interface {
	CrudDAO[models.User]
}
