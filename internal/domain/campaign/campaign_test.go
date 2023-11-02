package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {

	assert := assert.New(t)

	name := "Nome Campaign"
	content := "Conteúdo Content"
	contacts := []string{"email01@gmail.com", "email02@gmail.com"}

	campaign := NewCampaign(name, content, contacts)

	// if campaign.ID != "1" {
	// 	t.Errorf("Expected id 1, but got %s", campaign.ID)
	// } else if campaign.Name != name {
	// 	t.Errorf("Expected name %s, but got %s", name, campaign.Name)
	// } else if campaign.Content != content {
	// 	t.Errorf("Expected content %s, but got %s", content, campaign.Content)
	// } else if len(campaign.Contacts) != len(contacts) {
	// 	t.Errorf("Expected %d contacts, but got %d", len(contacts), len(campaign.Contacts))
	// } else if campaign.Contacts[0].Email != contacts[0] {
	// 	t.Errorf("Expected contact %s, but got %s", contacts[0], campaign.Contacts[0])
	// } else if campaign.Contacts[1].Email != contacts[1] {
	// 	t.Errorf("Expected contact %s, but got %s", contacts[1], campaign.Contacts[1])
	// }

	assert.Equal("1", campaign.ID)
	assert.Equal(name, campaign.Name)
	assert.Equal(content, campaign.Content)
	assert.Equal(len(contacts), len(campaign.Contacts))
	assert.Equal(contacts[0], campaign.Contacts[0].Email)
	assert.Equal(contacts[1], campaign.Contacts[1].Email)

}
