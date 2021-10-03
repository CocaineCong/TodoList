package api

import (
	"to-do-list/pkg/logging"
	"to-do-list/service"
	"github.com/gin-gonic/gin"
)

//TODO

// 创建商品
func CreateProduct(c *gin.Context) {
	service := service.CreateProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品列表
func ListTasks(c *gin.Context) {
	service := service.ListTasksService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.List()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func ListSameProducts(c *gin.Context)  {
	service := service.ListSameProductsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.ListSame()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

//商品详情
func ShowProduct(c *gin.Context) {
	service := service.ShowProductService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}

//删除商品
func DeleteProduct(c *gin.Context) {
	service := service.DeleteProductService{}
	res := service.Delete(c.Param("id"))
	c.JSON(200, res)
}

//更新商品
func UpdateProduct(c *gin.Context) {
	service := service.UpdateProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Update()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}

func UpProduct(c *gin.Context) {
	service := service.UpProductService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.UpProduct()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}


//搜索商品
func SearchProducts(c *gin.Context) {
	service := service.SearchProductsService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Show()
		c.JSON(200, res)
	} else {
		c.JSON(200, ErrorResponse(err))
		logging.Info(err)
	}
}
