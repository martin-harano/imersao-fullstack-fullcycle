package model

import (
	"encoding/json"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required"`
}

func (t *User) isValid() error {
	v := validator.New()
	err := v.Struct(t)
	if err != nil {
		fmt.Errorf("Error during User validation: %s", err.Error())
		return err
	}
	return nil
}

func (t *User) ParseJson(data []byte) error {
	err := json.Unmarshal(data, t)
	if err != nil {
		return err
	}

	err = t.isValid()
	if err != nil {
		return err
	}

	return nil
}

func (t *User) ToJson() ([]byte, error) {
	err := t.isValid()
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(t)
	if err != nil {
		return nil, nil
	}

	return result, nil
}

func NewUser() *User {
	return &User{}
}
