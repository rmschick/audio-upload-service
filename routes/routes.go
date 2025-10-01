package routes

import (
	"database/sql"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	upload "personal-dev/internal/handler"
	"personal-dev/internal/service/local"
)

// CORS sets up the CORS middleware.
func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost", "http://localhost:4200"},
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
	})
}

// SetupRouter initializes the Gin router with middleware and routes.
func SetupRouter(dbInstance *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(CORS())

	// Readiness probe; used for the Kubernetes deployment
	router.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	fileService := local.NewFileService(
		local.NewPostgresRepository(dbInstance),
		local.NewLocalStorer(),
	)

	uploadGroup := router.Group("/upload")
	{
		uploadGroup.POST("/", func(c *gin.Context) {
			upload.CreateUploadHandler(c, fileService)
		})

		uploadGroup.GET("/:id", func(c *gin.Context) {
			upload.GetUploadHandler(c, fileService)
		})
	}
	
	return router
}
