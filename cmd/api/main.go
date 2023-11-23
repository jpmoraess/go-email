package main

import (
	"go-email/internal/domain/campaign"
	"go-email/internal/dto"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.CampaignService{}

	r.Post("/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var input dto.NewCampaignInput
		render.DecodeJSON(r.Body, &input)
		id, err := campaignService.Create(input)
		if err != nil {
			render.Status(r, 400)
			render.JSON(w, r, map[string]string{"error": err.Error()})
			return
		}
		render.Status(r, 201)
		render.JSON(w, r, map[string]string{"id": id})
	})

	http.ListenAndServe(":8080", r)
}
