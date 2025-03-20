package application

import "github.com/BryanChanona/api_consumer/src/Responsible/domain"

type RegisterResponsibleUseCase struct {
	db domain.IResponsible
}


func NewRegisterResponsibleUseCase(db domain.IResponsible) *RegisterResponsibleUseCase {
	return &RegisterResponsibleUseCase{db: db}
}

func (useCase *RegisterResponsibleUseCase) Execute(responsible domain.Responsible) error{
	return useCase.db.SaveResponsible(responsible)
}