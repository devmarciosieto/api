package endpoints

import (
	"context"
	internalmock "github.com/devmarciosieto/api/internal/test/internal-mock"
	"github.com/go-chi/chi/v5"
	"net/http"
	"net/http/httptest"
)

var (
	service *internalmock.CampaignServerMock
	handler = Handler{}
)

func init() {
	service = new(internalmock.CampaignServerMock)
	handler.CampaignService = service
}

func newReqAndRecord(method string, url string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(method, url, nil)
	rr := httptest.NewRecorder()
	return req, rr
}

func addParams(req *http.Request, params map[string]string) *http.Request {
	chiContext := chi.NewRouteContext()
	for k, v := range params {
		chiContext.URLParams.Add(k, v)
	}
	return req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, chiContext))
}
