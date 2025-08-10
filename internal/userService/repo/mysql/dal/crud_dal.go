package dal

import (
	"context"
	"reflect"

	"gorm.io/gorm"
)

type crudDAL[T any] struct {
	DB *gorm.DB
}

func (u *crudDAL[T]) Create(entity *T) (*T, error) {
	ctx := context.Background()

	err := gorm.G[T](u.DB).Create(ctx, entity)

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *crudDAL[T]) GetByID(id int64) (*T, error) {
	var entity T

	if err := u.DB.First(&entity, id).Error; err != nil {
		return nil, err
	}

	return &entity, nil
}

func (u *crudDAL[T]) GetAll() ([]*T, error) {
	var entities []*T

	if err := u.DB.Limit(10).Offset(0).Find(&entities).Error; err != nil {
		return nil, err
	}

	return entities, nil
}

func (u *crudDAL[T]) Update(entity *T, id int64) (*T, error) {
	existingEntity, err := u.GetByID(id)

	if err != nil {
		return nil, err
	}

	updateData(entity, existingEntity)

	result := u.DB.Save(&existingEntity)

	if result.Error != nil {
		return nil, result.Error
	}

	return existingEntity, nil
}

func (u *crudDAL[T]) Delete(id int64) (bool, error) {
	var entity T

	if err := u.DB.First(&entity, id).Error; err != nil {
		return false, err
	}

	u.DB.Delete(&entity)

	return true, nil
}

func updateData[T any](entity *T, existingEntity *T) *T {
	valSrc := reflect.ValueOf(entity).Elem()
	valDst := reflect.ValueOf(existingEntity).Elem()

	for i := 0; i < valSrc.NumField(); i++ {
		srcField := valSrc.Field(i)
		dstField := valDst.Field(i)

		if srcField.IsValid() && !srcField.IsZero() && dstField.CanSet() {
			dstField.Set(srcField)
		}
	}

	return existingEntity
}
