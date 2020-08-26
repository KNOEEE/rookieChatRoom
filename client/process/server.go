package process
import (
	"fmt"
	"os"
	"go_code/projectLearn/chatRoom/client/utils"
	"go_code/projectLearn/chatRoom/common/message"
	"net"
	"encoding/json"
)

//显示登陆成功后的界面
func ShowMenu() {
	fmt.Println("------恭喜登陆成功------")
	fmt.Println("1.显示在线用户列表")
	fmt.Println("2.发送消息")
	fmt.Println("3.信息列表")
	fmt.Println("4.退出系统")
	fmt.Println("请选择（1-4）：")
	var key int
	var content string
	smsProcess := &SmsProcess{}
	fmt.Scanf("%d\n", &key)
	switch key {
	case 1:
		outputOnlineUsers()
	case 2:
		fmt.Println("请输入你要发送的内容：")
		//这里不能输入空格 不然就炸了
		fmt.Scanf("%s\n", &content)
		smsProcess.SendGroupMsg(content)
	case 3:
		fmt.Println("3.信息列表")
	case 4:
		fmt.Println("已退出系统")
		os.Exit(0)
	default:
		fmt.Println("指令错误，请输入1-4：")
	}
}


func serverProcessMsg(conn net.Conn) {
	//创建一个transfer实例 读取server发送的消息
	tf := &utils.Transfer{
		Conn: conn,
	}
	for {
		fmt.Printf("客户端正在等待服务器的消息\n")
		msg, err := tf.ReadPkg()
		if err != nil {
			fmt.Println("tf.ReadPkg() err =", err)
			return
		}
		switch msg.Type {
		case message.NotifyUserStatusMsgType:
			//exec
			var notifyStatusMsg message.NotifyUserStatusMsg
			json.Unmarshal([]byte(msg.Data), 
				&notifyStatusMsg)
			updateUserStatus(&notifyStatusMsg)
		case message.SmsMsgType:
			outputGroupMsg(&msg)
		default:
			fmt.Println("服务器端发送了未知类型的数据")
		}
	}
}