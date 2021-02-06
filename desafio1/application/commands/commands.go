package commands

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio1/application/usecase"
	"github.com/martin-harano/imersao-fullstack-fullcycle/desafio1/infrastructure/repository"
)

func TestAddUser(database *gorm.DB, name string, email string) {
	userRepository := repository.UserRepositoryDb{Db: database}
	userUseCase := usecase.UserUseCase{UserRepository: userRepository}

	user, err := userUseCase.AddNewUser(name, email)
	if err != nil {
		log.Fatal("cannot add user", err)
	}

	log.Printf("User successful added with [ID = %s, name = %s, e-mail = %s", user.ID, user.Name, user.Email)

}
