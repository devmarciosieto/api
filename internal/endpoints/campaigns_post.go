package endpoints

import (
	"github.com/devmarciosieto/api/internal/contract"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaignDto
	err := render.DecodeJSON(r.Body, &request)
	id, err := h.CampaignService.Create(request)
	return map[string]string{"id": id, "message": "campaign created"}, http.StatusCreated, err
}
