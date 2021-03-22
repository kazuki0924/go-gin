package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/kazuki0924/go-gin/api"
	"github.com/kazuki0924/go-gin/controller"
	"github.com/kazuki0924/go-gin/docs"
	"github.com/kazuki0924/go-gin/middlewares"
	"github.com/kazuki0924/go-gin/repository"
	"github.com/kazuki0924/go-gin/service"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	videoRepository repository.VideoRepository = repository.NewVideoRepository()

	videoService service.VideoService = service.New(videoRepository)
	loginService service.LoginService = service.NewLoginService()
	jwtService   service.JWTService   = service.NewJWTService()

	videoController controller.VideoController = controller.New(videoService)
	loginController controller.LoginController = controller.NewLoginController(loginService, jwtService)
)

// @securityDefinitions.apikey bearerAuth
// @in header
// @name Authorization
func main() {
	// Swagger 2.0 Meta Information
	docs.SwaggerInfo.Title = "Example gin API"
	docs.SwaggerInfo.Description = "example gin api"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	defer videoRepository.CloseDB()
	server := gin.Default()

	videoAPI := api.NewVideoAPI(loginController, videoController)

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{
		login := apiRoutes.Group("/auth")
		{
			login.POST("/token", videoAPI.Authenticate)
		}

		videos := apiRoutes.Group("/videos", middlewares.AuthorizeJWT())
		{
			videos.GET("", videoAPI.GetVideos)
			videos.POST("", videoAPI.CreateVideo)
			videos.PUT(":id", videoAPI.UpdateVideo)
			videos.DELETE(":id", videoAPI.DeleteVideo)
		}
	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// We can setup this env variable from the EB console
	port := os.Getenv("PORT")

	// Elastic Beanstalk forwards requests to port 5000
	if port == "" {
		port = "5000"
	}
	server.Run(":" + port)

}
