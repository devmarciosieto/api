package campaign

import (
	"errors"
	"github.com/devmarciosieto/api/internal/contract"
	internalerros "github.com/devmarciosieto/api/internal/internal-erros"
)

type Service interface {
	Create(newCampaign contract.NewCampaignDto) (string, error)
	GetBy(id string) (*contract.CampaignResponse, error)
	Cancel(id string) error
	Delete(id string) error
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(newCampaign contract.NewCampaignDto) (string, error) {

	campaign, err := NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails)
	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaign)

	if err != nil {
		return "", internalerros.ErrInternal
	}

	return campaign.ID, nil
}

func (s *ServiceImp) GetBy(id string) (*contract.CampaignResponse, error) {

	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, internalerros.ProcessErrorToReturn(err)
	}

	return &contract.CampaignResponse{
		ID:                   campaign.ID,
		Name:                 campaign.Name,
		Content:              campaign.Content,
		Status:               campaign.Status,
		AmountOfEmailsToSend: len(campaign.Contacts),
	}, nil

}

func (s *ServiceImp) Cancel(id string) error {

	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return internalerros.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("campaign status invalid")
	}

	campaign.Cancel()
	err = s.Repository.Update(campaign)

	if err != nil {
		return internalerros.ErrInternal
	}

	return nil
}

func (s *ServiceImp) Delete(id string) error {

	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return internalerros.ProcessErrorToReturn(err)
	}

	if campaign.Status != Pending {
		return errors.New("campaign status invalid")
	}

	campaign.Delete()
	err = s.Repository.Delete(campaign)

	if err != nil {
		return internalerros.ErrInternal
	}

	return nil
}
