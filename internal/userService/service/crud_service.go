package service

import (
	"userservice/internal/userService/repo/mysql/dao"
)

type CrudService[T any] struct {
	dao dao.CrudDAO[T]
}

func NewCrudService[T any](dao dao.CrudDAO[T]) *CrudService[T] {
	return &CrudService[T]{dao: dao}
}

func (s *CrudService[T]) RegisterEntity(entity *T) (*T, error) {
	return s.dao.Create(entity)
}

func (s *CrudService[T]) FetchEntities() ([]*T, error) {
	return s.dao.GetAll()
}

func (s *CrudService[T]) FetchEntity(id int64) (*T, error) {
	return s.dao.GetByID(id)
}

func (s *CrudService[T]) PatchEntity(entity *T, id int64) (*T, error) {
	return s.dao.Update(entity, id)
}

func (s *CrudService[T]) DeleteEntity(id int64) (bool, error) {
	return s.dao.Delete(id)
}
