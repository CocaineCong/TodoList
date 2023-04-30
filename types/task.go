package types

type ShowTaskReq struct {
}

type DeleteTaskReq struct {
}

type UpdateTaskReq struct {
	ID      uint   `form:"id" json:"id"`
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

type CreateTaskReq struct {
	Title   string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Content string `form:"content" json:"content" binding:"max=1000"`
	Status  int    `form:"status" json:"status"` // 0 待办   1已完成
}

type SearchTaskReq struct {
	Info string `form:"info" json:"info"`
}

type ListTasksReq struct {
	Limit int `form:"limit" json:"limit"`
	Start int `form:"start" json:"start"`
}
