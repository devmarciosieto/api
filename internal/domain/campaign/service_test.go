package campaign_test

import (
	"errors"
	"github.com/devmarciosieto/api/internal/domain/campaign"
	"gorm.io/gorm"
	"testing"

	"github.com/devmarciosieto/api/internal/contract"
	internalerros "github.com/devmarciosieto/api/internal/internal-erros"
	internalmock "github.com/devmarciosieto/api/internal/test/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaignDto{
		Name:      "Test Name",
		Content:   "Test body",
		Emails:    []string{"email1@gmail.com", "email2@gmail.com"},
		CreatedBy: "email@gmail.com",
	}
	repositoryMock *internalmock.CampaingMockRepository
	service        = campaign.ServiceImp{}
)

func setUp() {
	repositoryMock = new(internalmock.CampaingMockRepository)
	service.Repository = repositoryMock
}

func Test_Create_Save_Campaign(t *testing.T) {
	setUp()
	assert := assert.New(t)

	repositoryMock.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {

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
	setUp()
	newCampaign = contract.NewCampaignDto{
		Name:      "Test Name",
		Content:   "Test body",
		Emails:    []string{"email1@gmail.com", "email2@gmail.com"},
		CreatedBy: "email@gmail.com",
	}

	assert := assert.New(t)

	repositoryMock.On("Create", mock.Anything).Return(errors.New("internal server error"))

	_, err := service.Create(newCampaign)

	assert.NotNil(err)
	assert.Equal("internal server error", err.Error())

}

func Test_GetById_ReturnCampaign(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, "email@gmail.com")
	repositoryMock.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == campaign.ID
	})).Return(campaign, nil)

	campaignReturned, _ := service.GetBy(campaign.ID)

	assert.Equal(campaign.ID, campaignReturned.ID)
	assert.Equal(campaign.Name, campaignReturned.Name)
	assert.Equal(campaign.Content, campaignReturned.Content)
	assert.Equal(campaign.Status, campaignReturned.Status)
}

func Test_GetById_ReturnErrorWhenSomethingWrongExist(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaign, _ := campaign.NewCampaign(newCampaign.Name, newCampaign.Content, newCampaign.Emails, "email@gmail.com")
	repositoryMock.On("GetById", mock.Anything).Return(nil, errors.New("Something wrong"))

	_, err := service.GetBy(campaign.ID)

	assert.Equal(internalerros.ErrInternal.Error(), err.Error())

}

func Test_Delete_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignIdInvalid := "invalid"
	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.Delete(campaignIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_Delete_ReturnStatusInvalid_when_campaign_has_status_not_equals_pending(t *testing.T) {
	setUp()

	assert := assert.New(t)
	campaign := &campaign.Campaign{ID: "1", Status: campaign.Started}
	repositoryMock.On("GetById", mock.Anything).Return(campaign, nil)

	err := service.Delete(campaign.ID)

	assert.Equal("campaign status invalid", err.Error())
}

func Test_Delete_ReturnInternalError_when_delete_has_problem(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignFound, _ := campaign.NewCampaign("Test Name", "Test body", []string{"email@gmail.com"}, "email@gmail.com")
	repositoryMock.On("GetById", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(errors.New("error to delete campaign"))

	err := service.Delete(campaignFound.ID)

	assert.Equal(internalerros.ErrInternal.Error(), err.Error())
}

func Test_Delete_ReturnNil_when_delete_has_success(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignFound, _ := campaign.NewCampaign("Test Name", "Test body", []string{"email@gmail.com"}, "email@gmail.com")
	repositoryMock.On("GetById", mock.Anything).Return(campaignFound, nil)
	repositoryMock.On("Delete", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
		return campaignFound == campaign
	})).Return(nil)

	err := service.Delete(campaignFound.ID)

	assert.Nil(err)
}

func Test_StartCampaign_ReturnRecordNotFound_when_campaign_does_not_exist(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignIdInvalid := "invalid"
	repositoryMock.On("GetById", mock.Anything).Return(nil, gorm.ErrRecordNotFound)

	err := service.StartCampaign(campaignIdInvalid)

	assert.Equal(err.Error(), gorm.ErrRecordNotFound.Error())
}

func Test_StartCampaign_ReturnStatusInvalid_when_campaign_has_status_not_equals_pending(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaign := &campaign.Campaign{ID: "1", Status: campaign.Started}
	repositoryMock.On("GetById", mock.Anything).Return(campaign, nil)

	err := service.StartCampaign(campaign.ID)

	assert.Equal("campaign status invalid", err.Error())
}

func Test_StartCampaign_should_send_email(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignSaved := &campaign.Campaign{ID: "1", Status: campaign.Pending}
	repositoryMock.On("GetById", mock.Anything).Return(campaignSaved, nil)

	repositoryMock.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return campaignSaved.ID == campaignToUpdate.ID && campaignToUpdate.Status == campaign.Done
	})).Return(nil)

	emailWasSent := false
	sendEmail := func(campaign *campaign.Campaign) error {
		if campaign.ID == campaignSaved.ID {
			emailWasSent = true
		}
		return nil
	}

	service.SendEmail = sendEmail

	service.StartCampaign(campaignSaved.ID)

	assert.True(emailWasSent)

}

func Test_StartCampaign_ReturnError_when_func_SendMail_fail(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignSaved := &campaign.Campaign{ID: "1", Status: campaign.Pending}
	repositoryMock.On("GetById", mock.Anything).Return(campaignSaved, nil)
	sendEmail := func(campaign *campaign.Campaign) error {
		return errors.New("error to send email")
	}

	service.SendEmail = sendEmail

	err := service.StartCampaign(campaignSaved.ID)

	assert.Equal(internalerros.ErrInternal.Error(), err.Error())

}

func Test_StartCampaign_ReturnNil_when_updated_to_done(t *testing.T) {
	setUp()
	assert := assert.New(t)
	campaignSaved := &campaign.Campaign{ID: "1", Status: campaign.Pending}
	repositoryMock.On("GetById", mock.Anything).Return(campaignSaved, nil)

	repositoryMock.On("Update", mock.MatchedBy(func(campaignToUpdate *campaign.Campaign) bool {
		return campaignSaved.ID == campaignToUpdate.ID && campaignToUpdate.Status == campaign.Done
	})).Return(nil)

	sendEmail := func(campaign *campaign.Campaign) error {
		return nil
	}

	service.SendEmail = sendEmail

	service.StartCampaign(campaignSaved.ID)

	assert.Equal(campaign.Done, campaignSaved.Status)

}
