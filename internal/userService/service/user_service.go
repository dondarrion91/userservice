package service

import (
	"userservice/internal/userService/repo/mysql/dao"
	"userservice/pkg/models"
)

type UserService struct {
	dao dao.UserDAO
}

func NewUserService(dao dao.UserDAO) *UserService {
	return &UserService{dao: dao}
}

func (s *UserService) RegisterUser(user *models.User) (*models.User, error) {
	return s.dao.Create(user)
}

func (s *UserService) FetchUsers() ([]*models.User, error) {
	return s.dao.GetAll()
}

func (s *UserService) FetchUser(id int64) (*models.User, error) {
	return s.dao.GetByID(id)
}

func (s *UserService) PatchUser(user *models.User, id int64) (*models.User, error) {
	return s.dao.Update(user, id)
}

func (s *UserService) DeleteUser(id int64) (bool, error) {
	return s.dao.Delete(id)
}
