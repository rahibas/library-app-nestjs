package rest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"application/command"
	"domain/entities"
	"interface/api/rest"
	"interface/api/rest/dto/request"
	"interface/api/rest/dto/response"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateDealer(t *testing.T) {
	// Arrange
	mockService := NewMockDealerService()
	controller := rest.NewDealerController(echo.New(), mockService)

	// Create a dealer for testing
	dealer := entities.NewDealer("TestDealer")

	dealerJSON, _ := json.Marshal(dealer)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/dealers", bytes.NewReader(dealerJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// Act
	if err := controller.CreateDealerController(c); err != nil {
		t.Fatal(err)
	}

	fmt.Printf("rec: %s\n", rec.Body.String())

	// Assert
	assert.Equal(t, http.StatusCreated, rec.Code)

	var createdDealer entities.Dealer
	_ = json.Unmarshal(rec.Body.Bytes(), &createdDealer)
	assert.Equal(t, dealer.Name, createdDealer.Name)
}

func TestPutDealer(t *testing.T) {
	// Arrange
	mockService := NewMockDealerService()
	controller := rest.NewDealerController(echo.New(), mockService)

	createdDealer, err := mockService.CreateDealer(&command.CreateDealerCommand{Name: "TestDealer"})
	assert.NoError(t, err)

	updateRequest := request.UpdateDealerRequest{
		Id:   createdDealer.Result.Id,
		Name: "updatedName",
	}

	dealerJSON, _ := json.Marshal(updateRequest)
	req := httptest.NewRequest(http.MethodPut, "/api/v1/dealers", bytes.NewReader(dealerJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// Act
	if err := controller.PutDealerController(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var receivedResponse response.DealerResponse
	err = json.Unmarshal(rec.Body.Bytes(), &receivedResponse)
	assert.NoError(t, err)

	assert.Equal(t, updateRequest.Name, receivedResponse.Name)
}

func TestDeleteDealer(t *testing.T) {
	// Arrange
	mockService := NewMockDealerService()
	controller := rest.NewDealerController(echo.New(), mockService)

	createdDealer, err := mockService.CreateDealer(&command.CreateDealerCommand{Name: "TestDealer"})
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/api/v1/dealers/%s", createdDealer.Result.Id), nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// Act
	if err := controller.DeleteDealerController(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusNoContent, rec.Code)
}

func TestGetDealerById(t *testing.T) {
	// Arrange
	mockService := NewMockDealerService()
	controller := rest.NewDealerController(echo.New(), mockService)

	createdDealer, err := mockService.CreateDealer(&command.CreateDealerCommand{Name: "TestDealer"})
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/api/v1/dealers/%s", createdDealer.Result.Id), nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// Act
	if err := controller.GetDealerByIdController(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var fetchedDealer response.DealerResponse
	err = json.Unmarshal(rec.Body.Bytes(), &fetchedDealer)
	assert.NoError(t, err)

	assert.Equal(t, createdDealer.Result.Id.String(), fetchedDealer.Id)
	assert.Equal(t, createdDealer.Result.Name, fetchedDealer.Name)
}

func TestGetAllDealers(t *testing.T) {
	// Arrange
	mockService := NewMockDealerService()
	controller := rest.NewDealerController(echo.New(), mockService)

	_, err := mockService.CreateDealer(&command.CreateDealerCommand{Name: "TestDealer1"})
	assert.NoError(t, err)

	_, err = mockService.CreateDealer(&command.CreateDealerCommand{Name: "TestDealer2"})
	assert.NoError(t, err)

	req := httptest.NewRequest(http.MethodGet, "/api/v1/dealers", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)

	// Act
	if err := controller.GetAllDealersController(c); err != nil {
		t.Fatal(err)
	}

	// Assert
	assert.Equal(t, http.StatusOK, rec.Code)

	var dealers response.ListDealersResponse
	err = json.Unmarshal(rec.Body.Bytes(), &dealers)
	assert.NoError(t, err)

	assert.Equal(t, 2, len(dealers.Dealers))
}
