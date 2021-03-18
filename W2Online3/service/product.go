package service

import (
	"cmall/model"
	"cmall/pkg/e"
	"cmall/pkg/logging"
	"cmall/serializer"
	"fmt"
)

//展示商品详情的服务
type ShowProductService struct {
}

//删除商品的服务
type DeleteProductService struct {
}

//更新商品的服务
type UpdateProductService struct {
	ID            uint   `form:"id" json:"id"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	CategoryName  string  `form:"category_name" json:"category_name"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	Status 		  int   `form:"status" json:"status"`   //0 待办   1已完成
}

type UpProductService struct {
	ID            uint   `form:"id" json:"id"`
	CategoryID    int    `form:"category_id" json:"category_id"`
	CategoryName  string `form:"category_name" json:"category_name"`
	Title         string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info          string `form:"info" json:"info" binding:"max=1000"`
	Status 		  int    `form:"status" json:"status"`   //0 待办   1已完成
}



//搜索商品的服务
type SearchProductsService struct {
	Search string `form:"search" json:"search"`
}

// 商品
func (service *ShowProductService) Show(id string) serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	//增加点击数
	product.AddView()
	if product.CategoryID == 1 || product.CategoryID == 2 ||
		product.CategoryID == 3 || product.CategoryID == 4 ||
		product.CategoryID == 5 || product.CategoryID == 6 {
		product.AddTaskRank()
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

func (service *UpProductService) UpProduct() serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product,service.ID).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	product.Status = service.Status
	err = model.DB.Save(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProduct(product),
		Msg:    e.GetMsg(code),
	}
}

//删除商品
func (service *DeleteProductService) Delete(id string) serializer.Response {
	var product model.Product
	code := e.SUCCESS
	err := model.DB.First(&product, id).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	err = model.DB.Delete(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//更新商品
func (service *UpdateProductService) Update() serializer.Response {
	product := model.Product{
		CategoryName:	service.CategoryName,
		CategoryID:    service.CategoryID,
		Title:         service.Title,
		Info:          service.Info,
		Status:			service.Status,
	}
	product.ID = service.ID
	code := e.SUCCESS
	err := model.DB.Save(&product).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: code,
		Msg:    e.GetMsg(code),
	}
}

//搜索商品
func (service *SearchProductsService) Show() serializer.Response {
	products := []model.Product{}
	code := e.SUCCESS
	err := model.DB.Where("title LIKE ?", "%"+service.Search+"%").Find(&products).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	productsTemp := []model.Product{}
	err = model.DB.Where("info LIKE ? ",  "%"+service.Search+"%").Find(&productsTemp).Error
	if err != nil {
		logging.Info(err)
		code = e.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    e.GetMsg(code),
			Error:  err.Error(),
		}
	}
	products = append(products, productsTemp...)
	fmt.Println("products")
	fmt.Println(products)
	return serializer.Response{
		Status: code,
		Data:   serializer.BuildProducts(products),
		Msg:    e.GetMsg(code),
	}
}
