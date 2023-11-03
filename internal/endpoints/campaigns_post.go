package endpoints

import (
	"errors"
	"github.com/devmarciosieto/api/internal/contract"
	internalerros "github.com/devmarciosieto/api/internal/internal-erros"
	"github.com/go-chi/render"
	"log"
	"net/http"
)

func (h *Handler) CampaignPost(w http.ResponseWriter, r *http.Request) {
	var request contract.NewCampaignDto

	err := render.DecodeJSON(r.Body, &request)

	if err != nil {
		log.Println(err)
	}

	id, err := h.CampaignService.Create(request)

	if err != nil {

		if errors.Is(err, internalerros.ErrInternal) {
			render.Status(r, http.StatusInternalServerError)
		} else {
			render.Status(r, http.StatusBadRequest)
		}
		render.JSON(w, r, map[string]string{"error": err.Error()})
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, map[string]string{"id": id, "message": "campaign created"})
}
