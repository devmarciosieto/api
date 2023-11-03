package endpoints

import (
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {

	campaign, err := h.CampaignService.Repository.Get()
	return campaign, http.StatusOK, err
}
