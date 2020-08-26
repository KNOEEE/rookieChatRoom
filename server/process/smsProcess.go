package prcs
import (
	"fmt"
	"net"
	"go_code/projectLearn/chatRoom/common/message"
	"go_code/projectLearn/chatRoom/server/utils"
	"encoding/json"
)

//定义结构体
type SmsProcess struct {} //暂时不需要字段

func(this *SmsProcess) SendGroupMsg(msg *message.Message) {
	//遍历服务器端的map 发送信息
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	data, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("json.Marshal err =", err)
		return
	}
	for id, up := range userMgr.onlineUsers {
		//这里需要跳过自己 不能发给自己
		if id == smsMsg.UserId {
			continue
		}
		this.SendOnce(data, up.Conn)
	}
}

//服务端并不关心数据的内容 只管转发
func(this *SmsProcess) SendOnce(data []byte,
	conn net.Conn) {
	tf := &utils.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendOnce err =", err)
	}
}