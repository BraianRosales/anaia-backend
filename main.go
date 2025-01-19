package main

import (
	"anaia-backend/database"
	"anaia-backend/internal/repository"
	"anaia-backend/settings"
	"context"

	"github.com/jmoiron/sqlx"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		/*Registers one or more "constructors" that Fx can use to initialize dependencies.*/
		fx.Provide(
			context.Background,
			settings.New,
			database.New,
			repository.New,
		),

		/*Register a function that will be executed once all the necessary dependencies are ready.*/
		fx.Invoke(
			func(db *sqlx.DB) {
				_, err := db.Query("select * from USUARIOS")
				if err != nil {
					panic(err)
				}
			},
		),
	)

	app.Run()
}
