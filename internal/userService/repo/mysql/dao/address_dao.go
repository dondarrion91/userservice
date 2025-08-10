package dao

import "userservice/pkg/models"

type AddressDAO interface {
	CrudDAO[models.Address]
}
