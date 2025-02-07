package rest_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/sklinkert/go-ddd/internal/application/command"
	"github.com/sklinkert/go-ddd/internal/application/common"
	"github.com/sklinkert/go-ddd/internal/domain/entities"
	"github.com/sklinkert/go-ddd/internal/interface/api/rest/dto/response"
	"github.com/stretchr/testify/mock"

	"github.com/labstack/echo/v4"
	"github.com/sklinkert/go-ddd/internal/interface/api/rest"
	"github.com/stretchr/testify/assert"
)

func TestCreateVehicle(t *testing.T) {
	// Setup
	e := echo.New()
	mockService := new(MockVehicleService)
	reqBody := map[string]interface{}{"Name": "TestVehicle", "Price": 9.99, "DealerId": "123e4567-e89b-12d3-a456-426614174000"}
	reqBodyBytes, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/vehicles", bytes.NewReader(reqBodyBytes))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	ctrl := rest.NewVehicleController(e, mockService)

	createVehicleCommandResult := &command.CreateVehicleCommandResult{
		Result: &common.VehicleResult{
			Id:    uuid.New(),
			Name:  "TestVehicle",
			Price: 9.99,
		},
	}
	mockService.On("CreateVehicle", mock.Anything).Return(createVehicleCommandResult, nil)

	// Execute
	err := ctrl.CreateVehicleController(c)
	assert.NoError(t, err)

	// Deserialize the response body
	var responseBody map[string]interface{}
	err = json.Unmarshal(rec.Body.Bytes(), &responseBody)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	// Remove fields from responseBody that are not present in reqBody
	// For example, remove Id and Dealer fields
	delete(responseBody, "Id")
	delete(responseBody, "Dealer")
	delete(reqBody, "DealerId")
	delete(responseBody, "CreatedAt")
	delete(responseBody, "UpdatedAt")

	// Assertions
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.Equal(t, reqBody, responseBody)
	mockService.AssertExpectations(t)
}

func TestGetAllVehicles(t *testing.T) {
	// Setup
	e := echo.New()
	mockService := new(MockVehicleService) // Assuming you have a mock of VehicleService

	expectedVehicles := []*entities.Vehicle{
		{
			Id:    uuid.New(),
			Name:  "TestVehicle1",
			Price: 9.99,
		}, {
			Id:    uuid.New(),
			Name:  "TestVehicle2",
			Price: 14.99,
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/api/v1/vehicles", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := rest.NewVehicleController(e, mockService)
	mockService.On("FindAllVehicles").Return(expectedVehicles, nil)

	var expectedListResponse response.ListVehiclesResponse
	for _, vehicle := range expectedVehicles {
		expectedListResponse.Vehicles = append(expectedListResponse.Vehicles,
			&response.VehicleResponse{
				Id:    vehicle.Id.String(),
				Name:  vehicle.Name,
				Price: vehicle.Price,
			})
	}

	// Assertions
	if assert.NoError(t, ctrl.GetAllVehiclesController(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var receivedListResponse response.ListVehiclesResponse
		err := json.Unmarshal(rec.Body.Bytes(), &receivedListResponse)
		if assert.NoError(t, err) {
			assert.ElementsMatch(t, expectedListResponse.Vehicles, receivedListResponse.Vehicles)
		}
	}
}
