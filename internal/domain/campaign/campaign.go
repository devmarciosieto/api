package campaign

import (
	"time"

	internalerros "github.com/devmarciosieto/api/internal/internal-erros"

	"github.com/rs/xid"
)

const (
	Pending  = "Pending"
	Canceled = "Canceled"
	Delete   = "Delete"
	Started  = "Started"
	Done     = "Done"
	Fail     = "Fail"
)

type Contact struct {
	ID         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignID string `gorm:"size:50"`
}

type Campaign struct {
	ID        string    `gorm:"size:50"`
	Name      string    `validate:"min=5,max=50" gorm:"size:100"`
	CreatedOn time.Time `validate:"required"`
	Content   string    `validate:"required" gorm:"size:1024"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20"`
	CreatedBy string    `validate:"email" gorm:"size:100"`
}

// TODO: make unit test
func (c *Campaign) Start() {
	c.Status = Started
}

func (c *Campaign) Done() {
	c.Status = Done
}

// TODO: make unit test
func (c *Campaign) Fail() {
	c.Status = Fail
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
}

func (c *Campaign) Delete() {
	c.Status = Delete
}

func NewCampaign(name string, content string, emails []string, createdBy string) (*Campaign, error) {

	contacts := make([]Contact, len(emails))
	for index, email := range emails {
		contacts[index].ID = xid.New().String()
		contacts[index].Email = email
	}

	campaign := &Campaign{
		ID:        xid.New().String(),
		Name:      name,
		CreatedOn: time.Now(),
		Content:   content,
		Contacts:  contacts,
		Status:    Pending,
		CreatedBy: createdBy,
	}

	err := internalerros.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}
	return nil, err
}
