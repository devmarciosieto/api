package endpoints

import "github.com/devmarciosieto/api/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
