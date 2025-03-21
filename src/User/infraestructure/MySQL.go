package infraestructure

import (
	"database/sql"
	"log"

	"github.com/BryanChanona/api_consumer/src/User/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}

func (sql *MySQL) SaveUser(user domain.User) error {
	query, err := sql.db.Prepare("INSERT INTO `Usuario` (nombre,correo,password,premium) VALUES (?,?,?,?)")

	if err != nil {
		return err
	}
	defer query.Exec()

	_, err = query.Exec(user.Name, user.Email, user.Password, user.Premium)

	if err != nil {
		log.Println("Error saving book:", err)
		return err
	}
	return nil
}
