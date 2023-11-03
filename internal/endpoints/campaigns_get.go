package endpoints

import (
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {
	render.Status(r, http.StatusOK)
	render.JSON(w, r, h.CampaignService.Repository.Get())
}
