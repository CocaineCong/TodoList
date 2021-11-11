package serializer

type ResponseUser struct {
	Status int `json:"status" example:"200"`
	Data User `json:"data"`
	Msg string `json:"msg" example:"ok"`
	Error string `json:"error" example:""`
}

type ResponseTask struct {
	Status int `json:"status"`
	Data Task `json:"data"`
	Msg string `json:"msg"`
	Error string `json:"error"`
}
