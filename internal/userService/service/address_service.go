package service

import (
	"userservice/internal/userService/repo/mysql/dao"
	"userservice/pkg/models"
)

type AddressService struct {
	dao dao.AddressDAO
}

func NewAddressService(dao dao.AddressDAO) *AddressService {
	return &AddressService{dao: dao}
}

func (s *AddressService) RegisterAddress(address *models.Address) (*models.Address, error) {
	return s.dao.Create(address)
}

func (s *AddressService) FetchAddresses() ([]*models.Address, error) {
	return s.dao.GetAll()
}

func (s *AddressService) FetchAddress(id int64) (*models.Address, error) {
	return s.dao.GetByID(id)
}

func (s *AddressService) PatchAddress(address *models.Address, id int64) (*models.Address, error) {
	return s.dao.Update(address, id)
}

func (s *AddressService) DeleteAddress(id int64) (bool, error) {
	return s.dao.Delete(id)
}
