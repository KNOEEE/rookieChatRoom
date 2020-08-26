package process
import (
	"fmt"
	"go_code/projectLearn/chatRoom/common/message"
	"go_code/projectLearn/chatRoom/client/model"
)

//客户端要维护的map
var onlineUsers map[int]*message.User = make(map[int]*message.User,
	10)
var CurUser model.CurUser //在用户登陆成功后完成初始化

//显示当前在线用户
func outputOnlineUsers() {
	fmt.Println("当前在线用户列表：")
	for id, _ := range onlineUsers {
		fmt.Println("用户id：\t", id)
	}
}
func updateUserStatus(notifyStatusMsg *message.NotifyUserStatusMsg){
	user, ok := onlineUsers[notifyStatusMsg.UserId]
	if !ok {
		user = &message.User{
			UserId: notifyStatusMsg.UserId,
		}
	}
	user.UserStatus = notifyStatusMsg.Status
	onlineUsers[notifyStatusMsg.UserId] = user
	outputOnlineUsers()
}