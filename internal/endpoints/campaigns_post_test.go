package endpoints

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	internalmock "github.com/devmarciosieto/api/internal/test/mock"

	"github.com/devmarciosieto/api/internal/contract"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaignDto{
		Name:    "Campaign name",
		Content: "Content da campanha",
		Emails:  []string{"email@gmail.com"},
	}

	service := new(internalmock.CampaignServerMock)

	service.On("Create", mock.MatchedBy(func(request contract.NewCampaignDto) bool {
		return request.Name == body.Name &&
			request.Content == body.Content &&
			request.Emails[0] == body.Emails[0]

	})).Return("id", nil)

	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)
	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignPost(rr, req)

	assert.Equal(http.StatusCreated, status)
	assert.Nil(err)

}

func Test_CampaignsPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)

	body := contract.NewCampaignDto{
		Name:    "Campaign name",
		Content: "Content da campanha",
		Emails:  []string{"email@gmail.com"},
	}

	service := new(internalmock.CampaignServerMock)

	service.On("Create", mock.Anything).Return("", fmt.Errorf("error"))

	handler := Handler{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest("POST", "/campaigns", &buf)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignPost(rr, req)

	assert.NotNil(err)

}
