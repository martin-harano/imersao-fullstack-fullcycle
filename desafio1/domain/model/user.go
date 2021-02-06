package model

import (
	"errors"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

// UserRepositoryInteface is user interaction interface
type UserRepositoryInteface interface {
	AddUser(user *User) error
}

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

// User is a user
type User struct {
	Base  `valid:"required"`
	Name  string `json:"name" gorm:"type:varchar(20)" valid:"notnull"`
	Email string `json:"email" gorm:"type:varchar(20)" valid:"notnull"`
}

func (u *User) isValid() error {
	_, err := govalidator.ValidateStruct(u)

	if err != nil {
		return err
	}

	if u.Name == "" {
		return errors.New("Name cannot be empty")
	}

	if u.Email == "" {
		return errors.New("E-mail cannot be empty")
	}

	return nil
}

// NewUser creates a valid User
func NewUser(name string, email string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
