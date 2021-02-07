package factory

import (
	"github.com/jinzhu/gorm"
	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio2/application/usecase"
	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio2/infrastructure/repository"
)

func UserUseCaseFactory(database *gorm.DB) usecase.UserUseCase {
	userRepositoryDb := repository.UserRepositoryDb{Db: database}

	userUseCase := usecase.UserUseCase{
		UserRepository: &userRepositoryDb,
	}

	return userUseCase
}
