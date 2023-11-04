package campaign

import (
	"errors"
	"testing"

	"github.com/devmarciosieto/api/internal/contract"
	internalerros "github.com/devmarciosieto/api/internal/internal-erros"
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

func (m *MockRepository) Get() ([]Campaign, error) {
	//	args := m.Called(campaign)
	return nil, nil
}

func (m *MockRepository) GetBy(id string) (*Campaign, error) {
	//	args := m.Called(campaign)
	return nil, nil
}

var (
	newCampaign = contract.NewCampaignDto{
		Name:    "Test Name",
		Content: "Test body",
		Emails:  []string{"email1@gmail.com", "email2@gmail.com"},
	}
	repositoryMock = new(MockRepository)
	service        = ServiceImp{Repository: repositoryMock}
)

func Test_Create_Save_Campaign(t *testing.T) {

	assert := assert.New(t)

	repositoryMock.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {

		if campaign.ID == "" ||
			campaign.Name == "" ||
			campaign.Content == "" ||
			len(campaign.Contacts) == 0 ||
			campaign.Contacts[0].Email != newCampaign.Emails[0] ||
			campaign.Contacts[1].Email != newCampaign.Emails[1] {
			return false

		}

		return true
	})).Return(nil)

	service.Repository = repositoryMock

	id, err := service.Create(newCampaign)

	assert.NotNil(id)
	assert.Nil(err)

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {

	assert := assert.New(t)

	newCampaign.Name = ""
	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("name is required with min 5", err.Error())
	assert.False(errors.Is(internalerros.ErrInternal, err))
}

func Test_Create_ValidateRepositorySave(t *testing.T) {

	newCampaign = contract.NewCampaignDto{
		Name:    "Test Name",
		Content: "Test body",
		Emails:  []string{"email1@gmail.com", "email2@gmail.com"},
	}

	repositoryMock = new(MockRepository)
	service = ServiceImp{Repository: repositoryMock}

	assert := assert.New(t)

	repositoryMock.On("Save", mock.Anything).Return(errors.New("internal server error"))

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("internal server error", err.Error())

}
