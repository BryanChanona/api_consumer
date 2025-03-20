package domain

type IUser interface{
	SaveUser(user User) error
}