package routers

import "github.com/gin-gonic/gin"

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiv1 := r.Group("/api/v1")
	{
		CreateAssetRouters(apiv1)
		CreateTimeseriesRouters(apiv1)
	}

	return r
}
