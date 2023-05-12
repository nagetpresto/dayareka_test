package database

import (
	"BE/pkg/mysql"
	"fmt"
)

func RunMigration() {
	db := mysql.DB

	// Create users table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(255)
		)
	`)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	// Create transactions table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS transactions (
			id SERIAL PRIMARY KEY,
			user_id INT,
			menu VARCHAR(255),
			price INT,
			qty INT,
			total INT,
			payment VARCHAR(255),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id)
		)
	`)
	if err != nil {
		fmt.Println(err)
		panic("Migration Failed")
	}

	fmt.Println("Migration Success")
}
