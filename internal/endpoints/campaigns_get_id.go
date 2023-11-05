package endpoints

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func (h *Handler) CampaignGetId(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	id := chi.URLParam(r, "id")
	campaign, err := h.CampaignService.GetBy(id)

	return campaign, http.StatusOK, err
}
