package db

import (
	"database/sql"
	"encoding/json"
	"os"
	"errors"
	"github.com/TomDev24/GoSimpleService/internal/model"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
)

type Manager struct {
	db *sql.DB
}

func (d *Manager) Init() error {
	connStr := os.Getenv("DB_STR")
	if connStr == "" {
		return errors.New("Could not find environment variable")
	}
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return errors.New("Unable to connect to database")
	}
	d.db = db

	if err = d.CreateOrdersTable(); err != nil {
		return err
	}
	return nil
}

func (d *Manager) Close(){
	d.db.Close()
}

func (d *Manager) CreateOrdersTable() error {
	bytes, err := os.ReadFile("msc/schema.sql")
	if err != nil {
		return errors.New("Error while openening schema.sql")
	}

	_, err = d.db.Exec(string(bytes))
	if err != nil {
		return errors.New("Failed to create Orders table")
	}
	return nil
}

func (d *Manager) InsertOrder(id string, data []byte) error {
	_, err := d.db.Exec("INSERT INTO orders (id, data) VALUES($1, $2)", id, data)
	if err != nil {
		return err
	}
	return nil
}

func (d *Manager) GetAllOrders() ([]model.Order, error) {
	var order model.Order
	var orders []model.Order

	rows, err := d.db.Query("SELECT * FROM orders")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
        var id string
        var data string
        err = rows.Scan(&id, &data)
        if err != nil {
			return nil, err
        }
		err = json.Unmarshal([]byte(data), &order)
        if err != nil {
			return nil, err
        }
		orders = append(orders, order)
    }
	if err = rows.Err(); err != nil {
		return nil, err
    }
	return orders, nil
}
