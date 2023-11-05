package campaign

import (
	"github.com/devmarciosieto/api/internal/contract"
	internalerros "github.com/devmarciosieto/api/internal/internal-erros"
)

type Service interface {
	Create(newCampaign contract.NewCampaignDto) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaignDto) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)

	if err != nil {
		return "", internalerros.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {

	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, internalerros.ErrInternal
	}

	return &contract.CampaignResponse{
		ID:      campaign.ID,
		Name:    campaign.Name,
		Content: campaign.Content,
		Status:  campaign.Status,
	}, nil

}
