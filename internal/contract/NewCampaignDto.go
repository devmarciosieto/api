package contract

type NewCampaignDto struct {
	Name      string   `json:"name"`
	Content   string   `json:"content"`
	Emails    []string `json:"emails"`
	CreatedBy string   `json:"createdBy"`
}
