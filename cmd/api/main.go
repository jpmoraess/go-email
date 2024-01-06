package main

import (
	"go-email/internal/domain/campaign"
	"go-email/internal/endpoints"
	"go-email/internal/infra/database"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.CampaignService{
		CampaignRepository: &database.CampaignRepositoryAdapter{},
	}
	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", handler.CampaignPost)
	r.Get("/campaigns", handler.CampaignGet)

	http.ListenAndServe(":8080", r)
}
