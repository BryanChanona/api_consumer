package domain

type ITemperature interface {
	SaveTemperature(temperature Temperature) error
}