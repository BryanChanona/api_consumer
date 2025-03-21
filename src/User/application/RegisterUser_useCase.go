package application

import (
	"fmt"

	"github.com/BryanChanona/api_consumer/src/User/domain"
	"github.com/BryanChanona/api_consumer/src/helpers"
)

type RegisterUserUseCase struct {
	db domain.IUser
}

func NewRegisterUserUseCase(db domain.IUser) *RegisterUserUseCase{
	return &RegisterUserUseCase{db: db}
}

func (useCase *RegisterUserUseCase) Execute(user domain.User) error{
	password := user.Password 
	user.Premium = false

	
	hashPassword, err := helpers.EncryptPassword(password)

	if err != nil {
		fmt.Print("Hubo un error al hashear la contraseña.")
	}

	user.Password = string(hashPassword)

	return useCase.db.SaveUser(user)
}