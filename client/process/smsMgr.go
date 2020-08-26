package process
import (
	"fmt"
	"go_code/projectLearn/chatRoom/common/message"
	"encoding/json"
)

func outputGroupMsg(msg *message.Message) {
	//这里进来的类型确定是smsMsg类型的
	//反序列化
	var smsMsg message.SmsMsg
	err := json.Unmarshal([]byte(msg.Data), &smsMsg)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}

	//show
	info := fmt.Sprintf("用户id:%d\t对大家说:%s\n",
		smsMsg.UserId, smsMsg.Content)
	fmt.Println(info)
}