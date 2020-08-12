package routers

import (
	v1 "github.com/damingerdai/blog-service/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	article := v1.NewArticle()
	tag := v1.NewTag()
	apiV1 := r.Group("/api/v1")
	{
		apiV1.Any("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		apiV1.POST("/tag", tag.Create)
		apiV1.DELETE("/tag/:id", tag.Delete)
		apiV1.PUT("/tag/:id", tag.Update)
		apiV1.PATCH("/tag/:id/state", tag.Update)
		apiV1.GET("/tag/:id", tag.Get)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/article", article.Create)
		apiV1.DELETE("/article/:id", article.Delete)
		apiV1.PUT("/article/:id", article.Update)
		apiV1.PATCH("/article/:id/state", article.Update)
		apiV1.GET("/article/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}

	return r
}
