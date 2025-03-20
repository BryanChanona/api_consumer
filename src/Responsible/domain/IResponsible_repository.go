package domain

type IResponsible interface{
	SaveResponsible(responsible Responsible) error
}