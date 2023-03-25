package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vmoltaemcrkonrgcechd/online-store/config"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/routes"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
)

func Run(cfg *config.Config) error {
	app := fiber.New()

	app.Use(cors.New(cors.ConfigDefault))

	pq, err := pg.New(cfg)
	if err != nil {
		return err
	}

	routes.WithRouter(app,
		usecase.NewCityUseCase(repo.NewCityRepo(pq)),
	)

	return app.Listen(cfg.HTTPAddr)
}
