package http

import (
	"log"

	"github.com/go-chi/chi/v5"
	"github.com/sacredcrane/gitrocket/internal/infra/config"
)

type Server struct {
	config config.Config
	mux    chi.Mux
	logger log.Logger
}
