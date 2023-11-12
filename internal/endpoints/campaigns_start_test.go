package endpoints

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func Test_CampaignsStart_StatusOK(t *testing.T) {

	assert := assert.New(t)
	service.On("StartCampaign", mock.MatchedBy(func(id string) bool {
		return id == "123"
	})).Return(nil)
	req, rr := newReqAndRecord("PATCH", "/campaigns/start")
	req = addParams(req, map[string]string{"id": "123"})

	_, status, err := handler.CampaignStart(rr, req)
	assert.Equal(http.StatusOK, status)
	assert.Nil(err)

}

func Test_CampaignsStart_Error(t *testing.T) {
	assert := assert.New(t)
	errExpected := errors.New("error")
	service.On("StartCampaign", mock.Anything).Return(errExpected)
	req, rr := newReqAndRecord("PATCH", "/campaigns/start")

	_, _, err := handler.CampaignStart(rr, req)
	assert.Equal(errExpected, err)

}
