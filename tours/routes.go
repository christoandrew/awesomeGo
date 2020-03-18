package tours

import "github.com/gin-gonic/gin"

func GetTours(c *gin.RouterGroup)  {
	c.GET("/tours", _GetTours)
}