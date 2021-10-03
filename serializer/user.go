package serializer

import "to-do-list/model"

type User struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Status   string `json:"status"`
	CreateAt int64  `json:"create_at"`
}

//BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		ID:       user.ID,
		UserName: user.UserName,
		CreateAt: user.CreatedAt.Unix(),
	}
}

func BuildUsers(items []model.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
