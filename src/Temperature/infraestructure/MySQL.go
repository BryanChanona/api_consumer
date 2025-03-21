package infraestructure

import (
	
	"database/sql"
	"log"

	"github.com/BryanChanona/api_consumer/src/Temperature/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (mysql *MySQL) SaveTemperature(temperature domain.Temperature) error {

	query, err := mysql.db.Prepare("INSERT INTO registrotemperatura (medidaRegistrada) VALUES (?)")

	if err != nil {
		return err
	}
	defer query.Exec()

	_, err = query.Exec(temperature.RegisteredMeasure)

	if err != nil {
		log.Println("Error saving temperature", err)
		return err
	}
	return nil
}
