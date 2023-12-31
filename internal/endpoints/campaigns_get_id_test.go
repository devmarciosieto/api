package endpoints

import (
	"errors"
	"github.com/devmarciosieto/api/internal/contract"
	internalmock "github.com/devmarciosieto/api/internal/test/internal-mock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CampaignsGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)

	campaign := contract.CampaignResponse{
		ID:      "id",
		Name:    "Campaign name",
		Content: "Content da campanha",
		Status:  "Pending",
	}

	service := new(internalmock.CampaignServerMock)
	service.On("GetBy", mock.Anything).Return(&campaign, nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/campaigns", nil)
	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetId(rr, req)

	assert.Equal(http.StatusOK, status)
	assert.Equal(campaign.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaign.Name, response.(*contract.CampaignResponse).Name)

}

func Test_CampaignsGetById_should_return_error_when_something_wrong(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CampaignResponse{
		ID:      "id",
		Name:    "Campaign name",
		Content: "Content da campanha",
		Status:  "Pending",
	}

	service := new(internalmock.CampaignServerMock)
	errExpected := errors.New("something wrong")
	service.On("GetBy", mock.Anything).Return(&campaign, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/campaigns", nil)
	rr := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetId(rr, req)

	assert.Equal(errReturned.Error(), errExpected.Error())
}
