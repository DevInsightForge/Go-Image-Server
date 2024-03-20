package endpoint

import (
	"image-server/internal/application/service"

	"github.com/go-chi/chi/v5"
)

type ResizeEndpoint struct{}

func (rs ResizeEndpoint) Routes() chi.Router {
	router := chi.NewRouter()
	router.Get("/", service.ResizeHandler)

	return router
}
