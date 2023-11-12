package main

import (
	"github.com/devmarciosieto/api/internal/domain/campaign"
	"github.com/devmarciosieto/api/internal/infrastructure/database"
	"github.com/devmarciosieto/api/internal/infrastructure/email"
	"github.com/joho/godotenv"
)

func main() {

	println("Start World")

	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}

	db := database.NewDB()
	repository := database.CampaignRepository{Db: db}

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
		SendEmail:  email.SendEmail,
	}

	for {
		campaigns, err := repository.GetCampaignsToBeSent()

		if err != nil {
			println(err.Error())
		}

		println(len(campaigns))

		for _, campaign := range campaigns {
			campaignService.SendEmailAndUpdateStatus(&campaign)
		}

	}

}
