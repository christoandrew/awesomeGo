package products

import (
	"awesomeGo/common"
	"github.com/gin-gonic/gin"
	"net/http"
)
func Products(router *gin.RouterGroup) {
	router.GET("/products", _GetProducts)
}

func GetProduct(router *gin.RouterGroup) {
	router.GET("/products/:id", _GetProduct)
}

func _GetProducts(c *gin.Context) {
	db := common.Database()
	var products []Product
	db.Find(&products)
	c.JSON(http.StatusOK, products)
}

func _GetProduct(c *gin.Context) {
	db := common.Database()
	id := c.Param("id")
	var product Product
	db.First(&product).Where("id = ?", id)

	c.JSON(http.StatusOK, product)
}
