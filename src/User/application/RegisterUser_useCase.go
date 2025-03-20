package application

import "github.com/BryanChanona/api_consumer/src/User/domain"

type RegisterUserUseCase struct {
	db domain.IUser
}

func NewRegisterUserUseCase(db domain.IUser) *RegisterUserUseCase{
	return &RegisterUserUseCase{db: db}
}

func (useCase *RegisterUserUseCase) Execute(user domain.User) error{
	return useCase.db.SaveUser(user)
}