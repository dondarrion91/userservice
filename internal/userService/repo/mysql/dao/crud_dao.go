package dao

type CrudDAO[T any] interface {
	Create(entity *T) (*T, error)
	GetByID(id int64) (*T, error)
	GetAll() ([]*T, error)
	Update(entity *T, id int64) (*T, error)
	Delete(id int64) (bool, error)
}
