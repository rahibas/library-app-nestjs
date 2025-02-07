package rest

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sklinkert/go-ddd/internal/application/interfaces"
	"github.com/sklinkert/go-ddd/internal/interface/api/rest/dto/mapper"
	"github.com/sklinkert/go-ddd/internal/interface/api/rest/dto/request"
)

type DealerController struct {
	service interfaces.DealerService
}

func NewDealerController(e *echo.Echo, service interfaces.DealerService) *DealerController {
	controller := &DealerController{
		service: service,
	}

	e.POST("/api/v1/dealers", controller.CreateDealerController)
	e.GET("/api/v1/dealers", controller.GetAllDealersController)
	e.GET("/api/v1/dealers/:id", controller.GetDealerByIdController)
	e.PUT("/api/v1/dealers", controller.PutDealerController)
	e.DELETE("/api/v1/dealers/:id", controller.DeleteDealerController)

	return controller
}

func (sc *DealerController) CreateDealerController(c echo.Context) error {
	var createDealerRequest request.CreateDealerRequest

	if err := c.Bind(&createDealerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	dealerCommand, err := createDealerRequest.ToCreateDealerCommand()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid dealer Id format",
		})
	}

	commandResult, err := sc.service.CreateDealer(dealerCommand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create dealer",
		})
	}

	response := mapper.ToDealerResponse(commandResult.Result)

	return c.JSON(http.StatusCreated, response)
}

func (sc *DealerController) GetAllDealersController(c echo.Context) error {
	dealers, err := sc.service.FindAllDealers()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch dealers",
		})
	}

	response := mapper.ToDealerListResponse(dealers.Result)

	return c.JSON(http.StatusOK, response)
}

func (sc *DealerController) GetDealerByIdController(c echo.Context) error {
	// Hack: split the Id from the URL
	// For some reason c.Param("id") doesn't work here
	idRaw := c.Request().URL.Path[len("/api/v1/dealers/"):]

	id, err := uuid.Parse(idRaw)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid dealer Id format",
		})
	}

	dealer, err := sc.service.FindDealerById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch dealer",
		})
	}

	if dealer == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Dealer not found",
		})
	}

	response := mapper.ToDealerResponse(dealer.Result)

	return c.JSON(http.StatusOK, response)
}

func (sc *DealerController) PutDealerController(c echo.Context) error {
	var updateDealerRequest request.UpdateDealerRequest

	if err := c.Bind(&updateDealerRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	updateDealerCommand, err := updateDealerRequest.ToUpdateDealerCommand()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid dealer Id format",
		})
	}

	commandResult, err := sc.service.UpdateDealer(updateDealerCommand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to update dealer",
		})
	}

	response := mapper.ToDealerResponse(commandResult.Result)

	return c.JSON(http.StatusOK, response)
}

func (sc *DealerController) DeleteDealerController(c echo.Context) error {
	// Hack: split the Id from the URL
	// For some reason c.Param("id") doesn't work here
	idRaw := c.Request().URL.Path[len("/api/v1/dealers/"):]

	id, err := uuid.Parse(idRaw)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid dealer Id format",
		})
	}

	err = sc.service.DeleteDealer(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to delete dealer",
		})
	}

	return c.NoContent(http.StatusNoContent)
}
