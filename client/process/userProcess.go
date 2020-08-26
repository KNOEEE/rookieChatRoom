package process
import (
	"fmt"
	"net"
	"go_code/projectLearn/chatRoom/common/message"
	"go_code/projectLearn/chatRoom/client/utils"
	"encoding/binary"
	"encoding/json"
)

type UserProcess struct {

}

func (this *UserProcess) Register(userId int, userPwd string,
	userName string) (err error) {
	//连接server
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//通过conn发送消息
	var msg message.Message
	msg.Type = message.RegisterMsgType
	//创建一个结构体
	var registerMsg message.RegisterMsg
	registerMsg.User.UserId = userId
	registerMsg.User.UserPwd = userPwd
	registerMsg.User.UserName = userName

	//将msg序列化
	data, err := json.Marshal(registerMsg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	msg.Data = string(data)
	//将msg序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}

	tf := &utils.Transfer{
		Conn: conn,
	}
	//发送data给server
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("注册信息WritePkg err =", err)
		// return
	}
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err =", err)
		return
	}

	//反序列化
	var regResMsg message.RegisterResMsg
	err = json.Unmarshal([]byte(msg.Data), &regResMsg)
	if regResMsg.Code == 200 {
		fmt.Println("注册成功，请重新登陆")
	}else {
		fmt.Println(regResMsg.Error)
	}
	return
}

func (this *UserProcess) Login(userId int, userPwd string) (err error) {
	//连接server
	conn, err := net.Dial("tcp", "localhost:8888")
	if err != nil {
		fmt.Println("net.Dial err =", err)
		return
	}
	//延时关闭
	defer conn.Close()

	//通过conn发送消息
	var msg message.Message
	msg.Type = message.LoginMsgType
	//创建一个结构体
	var loginMsg message.LoginMsg
	loginMsg.UserId = userId
	loginMsg.UserPwd = userPwd
	data, err := json.Marshal(loginMsg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	msg.Data = string(data)

	//将msg序列化
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	//先发送data的长度
	//data的长度==》表示长度的[]byte
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(buf) err =", err)
		return
	}
	fmt.Println("客户端发送 消息长度", string(data))

	_, err = conn.Write(data)
	if err != nil {
		fmt.Println("conn.Write(data) err =", err)
		return
	}

	//处理返回的消息
	//创建一个实例
	tf := &utils.Transfer{
		Conn: conn,
	}
	msg, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("readPkg err =", err)
		return
	}
	//反序列化
	var loginResMsg message.LoginResMsg
	err = json.Unmarshal([]byte(msg.Data), &loginResMsg)
	if loginResMsg.Code == 200 {
		// 初始化CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//显示在线列表
		fmt.Println("当前在线用户列表如下：")
		for _, v := range loginResMsg.UsersId {
			//跳过自己
			if v == userId {
				continue
			}
			fmt.Println("用户ID:\t", v)
			//完成对onlineUsers的初始化
			user := &message.User{
				UserId: v,
				UserStatus: message.UserOnline,
			}
			onlineUsers[v] = user
		}
		fmt.Println()
		
		//启动一个协程保持和server的通信
		//如果server有数据推送给cli 可以在终端显示
		go serverProcessMsg(conn)
		for {
			ShowMenu()
		}
	}else {
		fmt.Println(loginResMsg.Error)
	}
	return
}