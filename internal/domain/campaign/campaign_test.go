package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name     = "Nome Campaign"
	content  = "Conteúdo Content"
	contacts = []string{"email01@gmail.com", "email02@gmail.com"}
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {

	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
	assert.Equal(contacts[0], campaign.Contacts[0].Email)
	assert.Equal(contacts[1], campaign.Contacts[1].Email)

	now := time.Now().Add(-time.Minute)
	assert.Greater(campaign.CreatedOn, now)
}

func Test_NewCampaign_InvalidName(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign("", content, contacts)

	assert.Equal("Name is required", err.Error())
}

func Test_NewCampaign_InvalidContacts(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("Contacts is required", err.Error())
}

func Test_NewCampaign_InvalidContent(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign(name, "", contacts)

	assert.Equal("Content is required", err.Error())
}

func Test_NewCampaign_InvalidEmail(t *testing.T) {

	assert := assert.New(t)

	_, err := NewCampaign(name, content, []string{})

	assert.Equal("Contacts is required", err.Error())
}
