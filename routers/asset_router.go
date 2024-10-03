package routers

import "github.com/gin-gonic/gin"

func CreateAssetRouters(r *gin.RouterGroup) error {

	r.GET("/assets", v1.GetAssets)
	r.POST("/assets", v1.AddAsset)
	r.PUT("/assets/:id", v1.EditAsset)
	r.DELETE("/assets/:id", v1.DeleteAsset)

	return nil
}
