package v1

import (
	"net/http"

	"effective/internal/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// Swagger docs.
	_ "github.com/oustrix/effective/docs"
)

// NewRouter -.
// Swagger spec:
//
//	@title			Effective Mobile
//	@description	Getting humans data
//	@version		1.0
//	@host			localhost:8080
//	@BasePath		/v1
func NewRouter(humansUC *usecase.HumansUseCase) http.Handler {
	r := gin.Default()

	// Swagger
	swaggerHandler := ginSwagger.WrapHandler(swaggerFiles.Handler)
	r.GET("/swagger/*any", swaggerHandler)

	// Routers
	h := r.Group("/v1")
	{
		v1Routes(h, humansUC)
	}

	return r
}

func v1Routes(r *gin.RouterGroup, humansUC *usecase.HumansUseCase) {
	setupHumansRoutes(r, humansUC)
}
