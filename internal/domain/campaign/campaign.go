package campaign

import (
	"time"

	internalerros "github.com/devmarciosieto/api/internal/internal-erros"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	ID        string    `validate:"required"`
	Name      string    `validate:"min=5,max=50"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"required"`
	Contacts  []Contact `validate:"min=1,dive"`
}

func NewCampaign(name string, content string, emails []string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
	}

	err := internalerros.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}
	return nil, err
}
