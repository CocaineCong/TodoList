package service

import (
	"to-do-list/model"
	"to-do-list/pkg/e"
	"to-do-list/pkg/logging"
	"to-do-list/serializer"
)

type CreateProductService struct {
	CategoryID   int    `form:"category_id" json:"category_id"`
	CategoryName string `form:"category_name" json:"category_name"`
	Title        string `form:"title" json:"title" bind:"required,min=2,max=100"`
	Info         string `form:"info" json:"info" bind:"max=1000"`
	Status       int   `form:"status" json:"status"`
	BossID       int    `form:"boss_id" json:"boss_id" bind:"required"`
	BossName     string `form:"boss_name" json:"boss_name"`
}
type ListProductsService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
	CategoryID uint `form:"category_id" json:"category_id"`
}

type ListSameProductsService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
	CategoryID uint `form:"category_id" json:"category_id"`

}

//创建标签
func (service *CreateProductService) Create() serializer.Response {
	product := model.Product{
		CategoryID:    service.CategoryID,
		CategoryName:  service.CategoryName,
		Title:         service.Title,
		Info:          service.Info,
		Status:          0,
		BossID:        service.BossID,
		BossName:      service.BossName,
	}
	code := e.SUCCESS
	err := model.DB.Create(&product).Error
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

func (service *ListProductsService) List() serializer.Response {
	products := []model.Product{}

	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}

	if service.CategoryID == 0 {
		if err := model.DB.Model(model.Product{}).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Limit(service.Limit).
			Offset(service.Start).Find(&products).
			Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := model.DB.Model(model.Product{}).
			Where("category_id=?", service.CategoryID).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("category_id=?", service.CategoryID).
			Limit(service.Limit).
			Offset(service.Start).
			Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}


func (service *ListSameProductsService) ListSame() serializer.Response {
	products := []model.Product{}

	total := 0
	code := e.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}

	if service.CategoryID == 0 {
		if err := model.DB.Model(model.Product{}).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Limit(service.Limit).
			Offset(service.Start).Find(&products).
			Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := model.DB.Model(model.Product{}).
			Where("category_id=?", service.CategoryID).
			Count(&total).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("category_id=?", service.CategoryID).
			Limit(service.Limit).
			Offset(service.Start).
			Find(&products).Error; err != nil {
			logging.Info(err)
			code = e.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    e.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}
