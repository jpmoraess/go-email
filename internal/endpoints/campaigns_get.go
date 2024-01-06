package endpoints

import (
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {
	render.Status(r, 201)
	render.JSON(w, r, h.CampaignService.CampaignRepository.Get())
}
