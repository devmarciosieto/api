package internal_mock

import (
	"github.com/devmarciosieto/api/internal/domain/campaign"
	"github.com/stretchr/testify/mock"
)

type CampaingMockRepository struct {
	mock.Mock
}

func (m *CampaingMockRepository) GetCampaignsToBeSent() ([]campaign.Campaign, error) {
	panic("implement me")
}

func (m *CampaingMockRepository) Create(campaign *campaign.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *CampaingMockRepository) Update(campaign *campaign.Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func (m *CampaingMockRepository) Get() ([]campaign.Campaign, error) {
	//	args := m.Called(campaign)
	return nil, nil
}

func (m *CampaingMockRepository) GetById(id string) (*campaign.Campaign, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*campaign.Campaign), nil
}

func (m *CampaingMockRepository) Delete(campaign *campaign.Campaign) error {
	args := m.Called(campaign)
	if args.Error(0) != nil {
		return args.Error(0)
	}
	return nil
}
