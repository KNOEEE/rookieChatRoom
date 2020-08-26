package process
import (
	"fmt"
	"go_code/projectLearn/chatRoom/common/message"
	"go_code/projectLearn/chatRoom/client/utils"
	"encoding/json"
)

type SmsProcess struct {}

//发送广播消息
func (this *SmsProcess) SendGroupMsg(content string) (err error){
	var msg message.Message
	msg.Type = message.SmsMsgType
	var smsMsg message.SmsMsg
	smsMsg.Content = content
	smsMsg.UserId = CurUser.UserId
	smsMsg.UserStatus = CurUser.UserStatus

	//序列化
	data, err := json.Marshal(smsMsg)
	if err != nil {
		fmt.Println("SendGroupMsg Marshal err", err)
		return
	}
	msg.Data = string(data)
	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("SendGroupMsg Marshal err", err)
		return
	}
	tf := &utils.Transfer{
		Conn: CurUser.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("SendGroupMsg WritePkg err", err)
		return
	}
	return
}