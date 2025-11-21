package services

import "github.com/go-chi/chi/v5"

type Mux interface {
	Router() *chi.Mux
}
