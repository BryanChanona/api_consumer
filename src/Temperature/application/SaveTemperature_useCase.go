package application


import (
	"fmt"

	"github.com/BryanChanona/api_consumer/src/Temperature/domain"
	
)

type SaveTemperatureUseCase struct {
	db domain.ITemperature
}

func NewSaveTemperatureUseCase(db domain.ITemperature) *SaveTemperatureUseCase{
	return &SaveTemperatureUseCase{db: db}
}

func (useCase *SaveTemperatureUseCase) Execute(temperature domain.Temperature) error{ 
	if temperature.RegisteredMeasure < 0 {
		return fmt.Errorf("temperatura no válida: %.2f. La temperatura no puede ser negativa", temperature.RegisteredMeasure)
	}


	return useCase.db.SaveTemperature(temperature)
}