package user

import (
	"ToDoListServer/orm"
)

type User struct {
	Name    string          `json:"name"`
	Pass    string          `json:"pass"`
}

// 这里做把数据从 Map 中获取出来的处理。
func Login(username string)string {
	password,bool := orm.DbmapServer.Get(username)
	if bool == true{
		return password
	}
	return nil
}
