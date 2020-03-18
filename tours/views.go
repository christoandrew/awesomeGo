package tours

import (
	"awesomeGo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func _GetTours(c *gin.Context) {
	db := common.Database()
	var tours []Tour
	db.Find(&tours)

	c.JSON(http.StatusOK, tours)
}
