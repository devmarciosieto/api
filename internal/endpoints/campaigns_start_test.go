package endpoints

import (
	"context"
	"errors"
	internalmock "github.com/devmarciosieto/api/internal/test/internal-mock"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_CampaignsStart_StatusOK(t *testing.T) {
	assert := assert.New(t)

	service := new(internalmock.CampaignServerMock)
	service.On("StartCampaign", mock.MatchedBy(func(id string) bool {
		return id == "123"
	})).Return(nil)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATCH", "/campaigns/start", nil)
	chiContext := chi.NewRouteContext()
	chiContext.URLParams.Add("id", "123")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))

	rr := httptest.NewRecorder()

	_, status, err := handler.CampaignStart(rr, req)
	assert.Equal(http.StatusOK, status)
	assert.Nil(err)

}

func Test_CampaignsStart_Error(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServerMock)
	errExpected := errors.New("error")
	service.On("StartCampaign", mock.Anything).Return(errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("PATCH", "/campaigns/start", nil)
	rr := httptest.NewRecorder()

	_, _, err := handler.CampaignStart(rr, req)
	assert.Equal(errExpected, err)

}
