package city_get

import (
	"context"
	"weatherapp/internal/domain"
)

func New() Handler {
	return Handler{}
}

type Handler struct {
}

func (h Handler) GetCity(ctx context.Context, id int64) (*domain.City, error) {
	city := domain.City{
		ID:   1,
		Name: "Moscow",
	}
	return &city, nil
}
