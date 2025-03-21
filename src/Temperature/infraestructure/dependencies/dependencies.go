package dependencies

import (
	"log"

	"github.com/BryanChanona/api_consumer/src/Temperature/application"
	"github.com/BryanChanona/api_consumer/src/Temperature/infraestructure"
	"github.com/BryanChanona/api_consumer/src/Temperature/infraestructure/controllers"
	"github.com/BryanChanona/api_consumer/src/helpers"
)

var (
	mySQL infraestructure.MySQL
)

func Init() {
	db, err := helpers.ConnMySQL()

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL =*infraestructure.NewMySQL(db)


}


func GetSaveTemperatureController() *controllers.SaveTemperatureController{
	useCase := application.NewSaveTemperatureUseCase(&mySQL)
	return controllers.NewSaveTemperatureController(useCase)

}