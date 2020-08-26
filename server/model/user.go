package model

//定义用户结构体
type User struct {
	//为了序列化 字符串的key与结构字段tag一致
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`
}