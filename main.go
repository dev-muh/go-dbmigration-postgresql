package main

import (
	"fmt"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/cast"
)

func main() {
	//// Start load file env ////
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Some error occured. Err: ", err.Error())
	}
	//// End load file env ////

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := cast.ToInt(os.Getenv("DB_PORT"))

	//// Start check connection database ////
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Println("Error Connecting DB => ", err)
		os.Exit(0)
	}
	defer db.Close()
	//// End check connection database ////

	//// Start Process DB Migration ////
	dbMigration := fmt.Sprintf("%s:%s@%s:%s/%s", username, password, host, cast.ToString(port), database)
	m, err := migrate.New("file://db-migrations/", "postgresql://"+dbMigration+"?sslmode=disable")

	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("--- Running Migration ---")
	m.Down()
	m.Up()
	fmt.Println("--- End Migration ---")
	//// Start Process DB Migration ////

	fmt.Println("Hello, Welcome in Service DB Migration")
}
