package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"log"
	"fmt"
)

type Manager struct {
	db *sql.DB
}

func (d *Manager) Init(){
	connStr := os.Getenv("DB_STR")
	//handle error
	db, err := sql.Open("pgx", connStr)
	handleError("Unable to connect to database", err)

	f, err := os.ReadFile("msc/schema.sql")
	handleError("Error while opening file", err)
	_, err = db.Exec(string(f))
	handleError("Failed to create table", err)
	d.db = db
}

func (d *Manager) InsertOrder(id string, data []byte){
	_, err := d.db.Exec("INSERT INTO orders (id, data) VALUES($1, $2)", id, data)
	if err != nil {
		fmt.Println(err)
	}
	//handleError("Could not insert", err)
}

func (d *Manager) ListAll() {
	rows, err := d.db.Query("SELECT * FROM orders")
	handleError("", err)
	defer rows.Close()
	for rows.Next() {
        var id string
        var data string
        err = rows.Scan(&id, &data)
        if err != nil {
            log.Fatalf("Failed to retrieve row because %s", err)
        }
		fmt.Println(id, data)
    }
	if err := rows.Err(); err != nil {
      log.Fatalf("Error encountered while iterating over rows: %s", err)
    }

}

func handleError(msg string, err error) {
    if err != nil {
        log.Fatalf("%s %s", msg, err)
	}
}

