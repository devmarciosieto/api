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

func (m *CampaignServerMock) GetBy(id string) (*contract.CampaignResponse, error) {
	args := m.Called(id)
	return args.Get(0).(*contract.CampaignResponse), args.Error(1)
}
