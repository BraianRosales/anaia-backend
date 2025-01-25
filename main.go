package main

import (
	"anaia-backend/database"
	"anaia-backend/internal/repository"
	"anaia-backend/internal/service"
	"anaia-backend/settings"
	"context"

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
			service.New,
		),

		/*Register a function that will be executed once all the necessary dependencies are ready.*/
		fx.Invoke(
			func(ctx context.Context, serv service.Service) {
				err := serv.RegisterUser(ctx, "Braian", "Rosales", "braianezequielrosales@gmail.com", "Braian154059")
				if err != nil {
					panic(err)
				}

				u, err := serv.LoginUser(ctx, "braianezequielrosales@gmail.com", "Braian154059")
				if err != nil {
					panic(err)
				}
				println(u.Name)
			},
		),
	)

	app.Run()
}
