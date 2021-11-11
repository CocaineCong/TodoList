package serializer

import (
	"to-do-list/model"
)

// swagger:response Resp
type Task struct {
	ID           uint      `json:"id" example:"1"`			// 任务ID
	Title        string    `json:"title" example:"吃饭"`	// 题目
	Content      string    `json:"content" example:"睡觉"`  // 内容
	View         uint64    `json:"view" example:"32"`  		// 浏览量
	Status       int 	   `json:"status" example:"0"`      // 状态(0未完成，1已完成)
	CreatedAt    int64     `json:"created_at"`
	StartTime 	 int64 `json:"start_time"`
	EndTime 	 int64 `json:"end_time"`
}

func BuildTask(item model.Task) Task {
	return Task{
		ID:           item.ID,
		Title:        item.Title,
		Content:      item.Content,
		Status:       item.Status,
		View:         item.View(),
		CreatedAt:    item.CreatedAt.Unix(),
		StartTime:    item.StartTime,
		EndTime:      item.EndTime,
	}
}

func BuildTasks(items []model.Task) (tasks []Task) {
	for _, item := range items {
		task := BuildTask(item)
		tasks = append(tasks, task)
	}
	return tasks
}
