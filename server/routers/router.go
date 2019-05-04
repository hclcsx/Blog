package routers

import (
	"gin-blog/server/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitializeRouter(r *gin.Engine)  {
	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}
}