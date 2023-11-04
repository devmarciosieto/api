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

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Post("/api/v1/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/api/v1/campaigns", endpoints.HandlerError(handler.CampaignGet))

	http.ListenAndServe(":8080", r)
}
