package main

import (
	"errors"
	"github.com/devmarciosieto/api/internal/infrastructure/database"
	internalerros "github.com/devmarciosieto/api/internal/internal-erros"
	"log"
	"net/http"

	"github.com/devmarciosieto/api/internal/contract"
	"github.com/devmarciosieto/api/internal/domain/campaign"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	service := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	r.Post("/api/v1/campaigns", func(w http.ResponseWriter, r *http.Request) {
		var request contract.NewCampaignDto

		err := render.DecodeJSON(r.Body, &request)

		if err != nil {
			log.Println(err)
		}

		id, err := service.Create(request)

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

	})

	http.ListenAndServe(":8080", r)
}
