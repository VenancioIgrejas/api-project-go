package v1

import (
	"github.com/VenancioIgrejas/api-project-go/pkg/app"
	"github.com/gin-gonic/gin"
)

func GetAssets(c *gin.Context) {
	appG := app.Gin{C: c}
	name := c.Query("name")
	state := -1
	if arg := c.Query("state"); arg != ""{
		state = com.
	}
}
