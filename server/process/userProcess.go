package prcs
import (
	"fmt"
	"net"
	"go_code/projectLearn/chatRoom/common/message"
	"go_code/projectLearn/chatRoom/server/utils"
	"go_code/projectLearn/chatRoom/server/model"
	"encoding/json"
)

type UserProcess struct {
	Conn net.Conn
	UserId int 
}

//通知所有在线users
func (this *UserProcess)NotifyOnlineUsers(userId int) {
	for id, up := range userMgr.onlineUsers{
		//不给自己发
		if id == userId {
			continue
		}
		up.NotifyOnline(userId)
	}
}

func(this *UserProcess) NotifyOnline(userId int) {
	var msg message.Message
	msg.Type = message.NotifyUserStatusMsgType
	var notifyStatusMsg message.NotifyUserStatusMsg
	notifyStatusMsg.UserId = userId
	notifyStatusMsg.Status = message.UserOnline

	//序列化
	data, err := json.Marshal(notifyStatusMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail", err)
		return
	}
	msg.Data = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Unmarshal fail", err)
		return
	}

	//创建transfer实例 
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyOnline fail", err)
		return
	}
}

func(this *UserProcess) ServerProcessReg(msg *message.Message) (err error) {
	//从msg中取出msg.data 反序列化
	var regMsg message.RegisterMsg
	err = json.Unmarshal([]byte(msg.Data), &regMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail", err)
		return
	}
	var resMsg message.Message
	resMsg.Type = message.RegisterResMsgType
	var regResMsg message.RegisterResMsg

	err = model.MyUserDao.Register(&regMsg.User)
	if err != nil {
		if err == model.ERROR_USER_EXISTS {
			regResMsg.Code = 505
			regResMsg.Error = model.ERROR_USER_EXISTS.Error()
		} else {
			regResMsg.Code = 506
			regResMsg.Error = "注册发生未知错误"
		}
	} else {
		regResMsg.Code = 200
	}

	//序列化 得到切片
	data, err := json.Marshal(regResMsg)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//将data赋值给resMsg
	resMsg.Data = string(data)
	//序列化
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//发送data
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}

//处理登陆请求
func (this *UserProcess) ServerProcessLogin(msg *message.Message) (err error) {
	//从msg中取出msg.data 反序列化
	var loginMsg message.LoginMsg
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("json.Unmarshal fail", err)
		return
	}

	var resMsg message.Message
	resMsg.Type = message.LoginResMsgType
	var loginResMsg message.LoginResMsg

	//到redis数据库完成验证
	user, err := model.MyUserDao.Login(loginMsg.UserId, 
		loginMsg.UserPwd)
	if err == nil {
		//legal
		loginResMsg.Code = 200
		//更新mgr
		//将userid赋值给this
		this.UserId = loginMsg.UserId
		userMgr.AddOnlineUser(this)
		this.NotifyOnlineUsers(loginMsg.UserId)

		for id, _ := range userMgr.onlineUsers {
			loginResMsg.UsersId = append(loginResMsg.UsersId,
				id)
		}
		fmt.Println(user, "login success")
	} else {
		//illegal
		if err == model.ERROR_USER_NOTEXISTS {
			loginResMsg.Code = 500 // user not exist
			loginResMsg.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMsg.Code = 403
			loginResMsg.Error = err.Error()
		} else {
			loginResMsg.Code = 505
			loginResMsg.Error = "服务器异常"
		}
	}
	
	//序列化 得到切片
	data, err := json.Marshal(loginResMsg)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}
	//将data赋值给resMsg
	resMsg.Data = string(data)
	//序列化
	data, err = json.Marshal(resMsg)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//发送data
	tf := &utils.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}