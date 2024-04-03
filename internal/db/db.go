package db

import (
	"database/sql"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"log"
	"fmt"
)

func Con() {
	//name := os.Getenv("POSTGRES_USER")
	connStr := os.Getenv("DB_STR")
	db, err := sql.Open("pgx", connStr)
    if err != nil {
        log.Fatalf("Unable to connect to database because %s", err)
    }
	f, err := os.ReadFile("msc/schema.sql")
	if err != nil {
        log.Fatalf("Error while opening file %s", err)
    }
	fmt.Println(string(f))
	_, err = db.Exec(string(f))
	if err != nil {
        log.Fatalf("Failed to create table%s", err)
    }

	rows, err := db.Query("SELECT * FROM orders")
	if err != nil {
        log.Fatalf("Database query failed because %s", err)
    }
	defer rows.Close()
	for rows.Next() {
        var id string
        var data string
        err = rows.Scan(&id, &data)
		fmt.Println(id, data)
        if err != nil {
            log.Fatalf("Failed to retrieve row because %s", err)
        }
    }

	if err := rows.Err(); err != nil {
      log.Fatalf("Error encountered while iterating over rows: %s", err)
    }
}
