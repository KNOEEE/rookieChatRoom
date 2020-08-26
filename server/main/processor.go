package main
import (
	"fmt"
	"net"
	"go_code/projectLearn/chatRoom/common/message"
	"go_code/projectLearn/chatRoom/server/utils"
	"go_code/projectLearn/chatRoom/server/process"
	"io"
	_ "encoding/json"
)

//定义结构体
type Processor struct {
	Conn net.Conn
}

//根据消息的类型不同 决定调用哪个函数来处理
func (this *Processor) serverProcessMsg(msg *message.Message) (err error) {
	fmt.Println("Msg =", msg)
	switch msg.Type {
	case message.LoginMsgType:
		//处理登陆逻辑
		//创建一个实例
		up := &prcs.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessLogin(msg)
	case message.RegisterMsgType:
		up := &prcs.UserProcess{
			Conn : this.Conn,
		}
		err = up.ServerProcessReg(msg)
	case message.SmsMsgType:
		sp := &prcs.SmsProcess{}
		sp.SendGroupMsg(msg)
	default:
		fmt.Println("消息类型错误，无法处理")
	}
	return
}

func (this *Processor) processPro() (err error) {
	//loop读信息
	for {
		//将读取数据包封装成函数readPkg
		//创建实例 读包
		tf := &utils.Transfer{
			Conn: this.Conn,
		}
		msg, err := tf.ReadPkg()
		if err != nil {
			if err == io.EOF {
				fmt.Println("cli exit, server exit too...")
				return err
			} else {
				fmt.Println("readPkg fail", err)
				return err
			}
			
		}
		err = this.serverProcessMsg(&msg)
		if err != nil {
			return err
		}
	}
}