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
	Name      string    `validate:"min=5,max=50" gorm:"size:100;not null"`
	CreatedOn time.Time `validate:"required" gorm:"not null"`
	UpdatedOn time.Time
	Content   string    `validate:"required" gorm:"size:1024;not null"`
	Contacts  []Contact `validate:"min=1,dive"`
	Status    string    `gorm:"size:20;not null"`
	CreatedBy string    `validate:"email" gorm:"size:100;not null"`
}

func (c *Campaign) Start() {
	c.Status = Started
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Done() {
	c.Status = Done
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Fail() {
	c.Status = Fail
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Cancel() {
	c.Status = Canceled
	c.UpdatedOn = time.Now()
}

func (c *Campaign) Delete() {
	c.Status = Delete
	c.UpdatedOn = time.Now()
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
