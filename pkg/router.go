package pkg

import "github.com/gofiber/fiber/v2"

type router struct {
	app        *fiber.App
	containers []Container
}

func NewRouter(app *fiber.App, containers []Container) Router {
	return &router{
		app:        app,
		containers: containers,
	}
}

func (r *router) Handle() {
	api := r.app.Group("/api")
	r.HandleApi(api)

	r.HandleWeb(r.app)
}

func (r *router) HandleApi(router fiber.Router) {
	for _, container := range r.containers {
		container.HandleApi(router)
	}
}

func (r *router) HandleWeb(router fiber.Router) {
	for _, container := range r.containers {
		container.HandleWeb(router)
	}
}
