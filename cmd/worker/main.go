package main

import (
	"github.com/devmarciosieto/api/internal/infrastructure/database"
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

	for {
		campaigns, err := repository.GetCampaignsToBeSent()

		if err != nil {
			println(err.Error())
		}

		for _, campaign := range campaigns {
			println(campaign.ID)
		}

	}

}
