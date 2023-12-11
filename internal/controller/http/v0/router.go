package v0

import (
	"github.com/Nikola-zim/3d-printing-studio/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(handler *gin.Engine, l zerolog.Logger, o usecase.OrderManager) {
	// Options
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	// Routers
	h := handler.Group("/v0")
	{
		newOrderManagerRouter(h, l, o)
	}
}
