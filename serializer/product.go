package serializer

import "to-do-list/model"

type Product struct {
	ID           uint   `json:"id"`
	CategoryID   int    `json:"category_id"`
	CategoryName  string    `json:"category_name"`
	Title        string `json:"title"`
	Info         string `json:"info"`
	View         uint64 `json:"view"`
	Status       int `json:"status"`
	CreatedAt    int64  `json:"created_at"`
	BossID       int    `json:"boss_id"`
	BossName     string `json:"boss_name"`
}

// 序列化商品
func BuildProduct(item model.Product) Product {
	return Product{
		ID:           item.ID,
		CategoryID:   item.CategoryID,
		CategoryName: item.CategoryName,
		Title:        item.Title,
		Info:         item.Info,
		Status:       item.Status,
		View:         item.View(),
		CreatedAt:    item.CreatedAt.Unix(),
		BossID:       item.BossID,
		BossName:     item.BossName,
	}
}

//序列化商品列表
func BuildProducts(items []model.Product) (products []Product) {
	for _, item := range items {
		product := BuildProduct(item)
		products = append(products, product)
	}
	return products
}
