package endpoints

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	internalerros "github.com/devmarciosieto/api/internal/internal-erros"
	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_when_endpoint_returns_internal_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerros.ErrInternal
	}

	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)

	assert.Equal(http.StatusInternalServerError, res.Code)
	assert.Contains(res.Body.String(), internalerros.ErrInternal.Error())

}

func Test_HandlerError_when_endpoint_returns_domain_error(t *testing.T) {
	assert := assert.New(t)

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, errors.New("domain error")
	}

	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)

	assert.Equal(http.StatusBadRequest, res.Code)
	assert.Contains(res.Body.String(), "domain error")

}

func Test_HandlerError_when_endpoint_returns_obj_and_status(t *testing.T) {
	assert := assert.New(t)
	type bodyForTest struct {
		Id int
	}

	objExpected := bodyForTest{Id: 2}

	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return objExpected, http.StatusCreated, nil
	}

	handlerFunc := HandlerError(endpoint)
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	handlerFunc(res, req)

	assert.Equal(http.StatusCreated, res.Code)
	objReturned := bodyForTest{}
	json.Unmarshal(res.Body.Bytes(), &objReturned)
	assert.Equal(objExpected, objReturned)

}
