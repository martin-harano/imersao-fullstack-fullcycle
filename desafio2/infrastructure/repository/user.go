package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio2/domain/model"
)

// UserRepositoryDb is
type UserRepositoryDb struct {
	Db *gorm.DB
}

func (r UserRepositoryDb) AddUser(user *model.User) error {
	err := r.Db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
