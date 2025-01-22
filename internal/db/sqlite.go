package db

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/shakibhasan09/imager/internal/db/models"
)

func init() {
	db, err := sqlx.Open("sqlite3", "../../database.db")
	if err != nil {
		log.Fatal(err)
	}

	tx := db.MustBegin()
	tx.NamedExec("INSERT INTO projects (name) VALUES (:name)", &models.Project{Name: "test"})
	tx.Commit()
}
