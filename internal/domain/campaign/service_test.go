package campaign

import (
	"testing"

	"github.com/devmarciosieto/api/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Save(campaign *Campaign) error {
	args := m.Called(campaign)
	return args.Error(0)
}

func Test_Create_Save_Campaign(t *testing.T) {

	assert := assert.New(t)

	newCampaign := contract.NewCampaignDto{
		Name:    "Test Name",
		Content: "Test body",
		Emails:  []string{"email1@gmail.com", "email2gmail.com"},
	}

	repositoryMock := new(MockRepository)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {

		if campaign.ID == "" {
			return false

		} else if campaign.Name != newCampaign.Name {
			return false
		} else if campaign.Content != newCampaign.Content {
			return false
		} else if len(campaign.Contacts) != len(newCampaign.Emails) {
			return false
		} else if campaign.Contacts[0].Email != newCampaign.Emails[0] {
			return false
		} else if campaign.Contacts[1].Email != newCampaign.Emails[1] {
			return false
		}
		return true
	})).Return(nil)

	service := Service{Repository: repositoryMock}

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)

	repositoryMock.AssertExpectations(t)
}
