package usecase

import "github.com/martin-harano/imersao-fullstack-fullcycle/desafio2/domain/model"

type UserUseCase struct {
	UserRepository model.UserRepositoryInteface
}

func (u *UserUseCase) AddNewUser(name string, email string) (*model.User, error) {

	user, err := model.NewUser(name, email)
	if err != nil {
		return nil, err
	}

	u.UserRepository.AddUser(user)
	if user.ID == "" {
		return nil, err
	}
	return user, nil
}
