package endpoints

import (
	"errors"
	"go-email/internal/dto"
	internalerrors "go-email/internal/internal-errors"
	"net/http"

	"github.com/go-chi/render"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var input dto.NewCampaignInput
	render.DecodeJSON(r.Body, &input)
	id, err := h.CampaignService.Create(input)
	if err != nil {
		if errors.Is(err, internalerrors.ErrInternal) {
			render.Status(r, 500)
		} else {
			render.Status(r, 400)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}
	render.Status(r, 201)
	render.JSON(w, r, map[string]string{"id": id})
}
