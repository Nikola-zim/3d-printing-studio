package v0

import (
	"github.com/Nikola-zim/3d-printing-studio/internal/entity"
	"github.com/Nikola-zim/3d-printing-studio/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
)

type orderManager struct {
	orderManager usecase.OrderManager
	log          zerolog.Logger
}

func newOrderManagerRouter(handler *gin.RouterGroup, log zerolog.Logger, o usecase.OrderManager) {
	r := &orderManager{o, log}

	h := handler.Group("/client-page")
	{
		h.GET("/my-orders", r.getOrders)
	}
}

type getOrdersResponse struct {
	History []entity.Order `json:"orders"`
}

func (om *orderManager) getOrders(c *gin.Context) {
	translations, err := om.orderManager.GetOrders(c.Request.Context(), 1)
	if err != nil {
		om.log.Info().Msgf("failed to get orders: %s", err)
		errorResponse(c, http.StatusInternalServerError, "get order problems")

		return
	}

	c.JSON(http.StatusOK, getOrdersResponse{translations})
}
