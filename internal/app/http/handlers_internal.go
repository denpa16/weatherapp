package http

import (
	"net/http"
	httpCommandCityGet "weatherapp/internal/services/http/city_get"
)

func RegisterInternalHandlers(mux *http.ServeMux) {
	// Get city
	mux.Handle(
		"GET /api/city/{id}/",
		NewCityGetHandler(
			httpCommandCityGet.New(),
			"GET /api/city/{id}/",
		),
	)
}
