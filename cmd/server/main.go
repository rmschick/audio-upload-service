package main

import (
	"os"

	"personal-dev/internal/db"
	"personal-dev/routes"
)

func main() {
	databaseInstance, err := db.InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		panic(err)
	}
	defer db.CloseDB(databaseInstance)

	if err := routes.SetupRouter(databaseInstance).Run(":8080"); err != nil {
		panic(err)
	}
}
