package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/vmoltaemcrkonrgcechd/online-store/config"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/routes"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase"
	"github.com/vmoltaemcrkonrgcechd/online-store/internal/usecase/repo"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/pg"
	"github.com/vmoltaemcrkonrgcechd/online-store/pkg/sessionstore"
)

func Run(cfg *config.Config) error {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:8080",
		AllowMethods:     "GET,POST,HEAD,DELETE,PATCH",
		AllowCredentials: true,
	}))

	pq, err := pg.New(cfg)
	if err != nil {
		return err
	}

	routes.WithRouter(app,
		usecase.NewCityUseCase(repo.NewCityRepo(pq)),
		usecase.NewImageUseCase(repo.NewImageRepo(pq), cfg.FileSystem),
		usecase.NewUserUseCase(repo.NewUserRepo(pq)),
		usecase.NewColorUseCase(repo.NewColorRepo(pq)),
		usecase.NewCategoryUseCase(repo.NewCategoryRepo(pq)),
		usecase.NewProductUseCase(repo.NewProductRepo(pq)),
		sessionstore.New(),
	)

	return app.Listen(cfg.HTTPAddr)
}
