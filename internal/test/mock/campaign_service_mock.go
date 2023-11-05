package mock

import (
	"github.com/devmarciosieto/api/internal/contract"
	"github.com/stretchr/testify/mock"
)

type CampaignServerMock struct {
	mock.Mock
}

func (m *CampaignServerMock) Create(newCampaign contract.NewCampaignDto) (string, error) {
	args := m.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (m *CampaignServerMock) Update(newCampaign contract.NewCampaignDto) (string, error) {
	args := m.Called(newCampaign)
	return args.String(0), args.Error(1)
}

func (m *CampaignServerMock) GetBy(id string) (*contract.CampaignResponse, error) {
	args := m.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*contract.CampaignResponse), args.Error(1)
}

func (m *CampaignServerMock) Cancel(id string) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *CampaignServerMock) Delete(id string) error {
	args := m.Called(id)
	if args.Error(0) != nil {
		return args.Error(0)
	}
	return nil
}
