package rest

import (
	"net/http"

	"application/interfaces"
	"interface/api/rest/dto/mapper"
	"interface/api/rest/dto/request"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type VehicleController struct {
	service interfaces.VehicleService
}

func NewVehicleController(e *echo.Echo, service interfaces.VehicleService) *VehicleController {
	controller := &VehicleController{
		service: service,
	}

	e.POST("/api/v1/vehicles", controller.CreateVehicleController)
	e.GET("/api/v1/vehicles", controller.GetAllVehiclesController)
	e.GET("/api/v1/vehicles/:id", controller.GetVehicleByIdController)
	e.Use(middleware.Recover())

	return controller
}

func (pc *VehicleController) CreateVehicleController(c echo.Context) error {
	var createVehicleRequest request.CreateVehicleRequest

	if err := c.Bind(&createVehicleRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	vehicleCommand, err := createVehicleRequest.ToCreateVehicleCommand()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid vehicle Id format",
		})
	}

	result, err := pc.service.CreateVehicle(vehicleCommand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create vehicle",
		})
	}

	response := mapper.ToVehicleResponse(result.Result)

	return c.JSON(http.StatusCreated, response)
}

func (pc *VehicleController) GetAllVehiclesController(c echo.Context) error {
	vehicles, err := pc.service.FindAllVehicles()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch vehicles",
		})
	}

	response := mapper.ToVehicleListResponse(vehicles.Result)

	return c.JSON(http.StatusOK, response)
}

func (pc *VehicleController) GetVehicleByIdController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid vehicle Id format",
		})
	}

	vehicle, err := pc.service.FindVehicleById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch vehicle",
		})
	}

	if vehicle == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Vehicle not found",
		})
	}

	response := mapper.ToVehicleResponse(vehicle.Result)

	return c.JSON(http.StatusOK, response)
}
