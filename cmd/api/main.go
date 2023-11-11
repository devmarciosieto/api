package main

import (
	"net/http"

	"github.com/devmarciosieto/api/internal/endpoints"
	"github.com/devmarciosieto/api/internal/infrastructure/database"

	"github.com/devmarciosieto/api/internal/domain/campaign"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)

	db := database.NewDB()

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}
	r.Get("/api/v1/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/api/v1/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetId))
		r.Patch("/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
	})

	http.ListenAndServe(":8081", r)
}
