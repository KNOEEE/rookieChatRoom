package message

//确定消息类型
const (
	LoginMsgType = "LoginMsg"
	LoginResMsgType = "LoginResMsg"
	RegisterMsgType = "RegisterMsg"
	RegisterResMsgType = "RegisterResMsg"
	NotifyUserStatusMsgType = "NotifyUserStatusMsg"
	SmsMsgType = "SmsMsg"
)

//定义几个用户状态常量
const (
	UserOnline = iota
	UserOffline
	UserBusy
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

//定义2个消息
type LoginMsg struct {
	UserId int `json:"userId"`
	UserPwd string `json:"userPwd"`
	UserName string `json:"userName"`//用户名
}

type LoginResMsg struct {
	Code int `json:"code"`// 状态码 500用户不存在 200OK
	Error string `json:"error"`// 返回的错误信息
	UsersId []int
}

type RegisterMsg struct {
	User User `json:"user"`
}

type RegisterResMsg struct {
	Code int `json:"code"`// 状态码 400用户被占用 200OK
	Error string `json:"error"`
}

//配合server推送用户状态变化消息
type NotifyUserStatusMsg struct {
	UserId int `json:"userId"`
	Status int `json:"status"`
}

type SmsMsg struct {
	Content string `json:"content"`
	User // 匿名结构体，继承
}