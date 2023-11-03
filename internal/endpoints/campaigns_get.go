package endpoints

import (
	"errors"
	internalerros "github.com/devmarciosieto/api/internal/internal-erros"
	"github.com/go-chi/render"
	"net/http"
)

func (h *Handler) CampaignGet(w http.ResponseWriter, r *http.Request) {

	campaign, err := h.CampaignService.Repository.Get()

	if err != nil {
		if errors.Is(err, internalerros.ErrInternal) {
			render.Status(r, http.StatusInternalServerError)
		} else {
			render.Status(r, http.StatusBadRequest)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaign)
}
