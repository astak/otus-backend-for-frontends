package router

import (
	"regexp"

	ginprometheus "github.com/Astak/go-gin-prometheus"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/config"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/handler"
	"github.com/Astak/otus-docker-basics-homework/web-service-gin/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *handler.Handler, cfg config.Config) *gin.Engine {
	engine := gin.Default()
	engine.GET("/health", handler.GetHealth)

	p := ginprometheus.NewPrometheus("gin")
	re := regexp.MustCompile(`\/?user\/(\d+)`)
	p.ReqCntURLLabelMappingFn = func(c *gin.Context) string {
		return re.ReplaceAllString(c.Request.URL.Path, "/user/:id")
	}
	p.Use(engine)

	api := engine.Group("/api")
	{
		secured := api.Use(middlewares.Auth(cfg.GetJwtSecret()))
		{
			secured.GET("/user/profile", handler.GetProfile)
			secured.PUT("/user/profile", handler.UpdateProfile)
		}
	}

	return engine
}
