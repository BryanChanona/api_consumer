package dependencies

import (
	"log"

	"github.com/BryanChanona/api_consumer/src/User/application"
	"github.com/BryanChanona/api_consumer/src/User/infraestructure"
	"github.com/BryanChanona/api_consumer/src/User/infraestructure/controllers"
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

func GetSaveUserController() *controllers.UserController{
	useCaseSaveUser := application.NewRegisterUserUseCase(&mySQL)
	return controllers.NewUserController(useCaseSaveUser)
}