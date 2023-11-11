package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "Nome Campaign"
	content   = "Conteúdo Content"
	contacts  = []string{"email01@gmail.com", "email02@gmail.com"}
	createdBy = "email@gmail.com"
	fake      = faker.New()
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.NotNil(campaign.ID)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
	assert.Equal(contacts[0], campaign.Contacts[0].Email)
	assert.Equal(contacts[1], campaign.Contacts[1].Email)
	assert.Equal(createdBy, campaign.CreatedBy)

	now := time.Now().Add(-time.Minute)
	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_InvalidNameMin(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts, createdBy)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_InvalidNameMax(t *testing.T) {

	assert := assert.New(t)
	_, err := NewCampaign(fake.Lorem().Text(80), content, contacts, createdBy)

	assert.Equal("name is required with max 50", err.Error())
}

func Test_NewCampaign_InvalidContacts(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{}, createdBy)

	assert.Equal("contacts is required with min 1", err.Error())
}

func Test_NewCampaign_InvalidContent(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts, createdBy)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_InvalidEmail(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{"email"}, createdBy)

	assert.Equal("email is not valid", err.Error())
}

func Test_NewCampaign_IDIsNotNill(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_MustStatusStartWithPending(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts, createdBy)

	assert.Equal(Pending, campaign.Status)
}
