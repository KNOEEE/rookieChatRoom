package main
import (
	"fmt"
	"os"
	"go_code/projectLearn/chatRoom/client/process"
)
//go build -o client.exe go_code\projectLearn\chatRoom\client

//定义2个变量 一个表示用户id 一个表示用户pwd
var userId int
var userPwd string
var userName string

func main() {
	//接收用户的选择
	var key int
	for {
		fmt.Println("-----------欢迎登陆多人聊天系统-----------")
		fmt.Println("\t\t 1 登陆聊天室")
		fmt.Println("\t\t 2 注册用户")
		fmt.Println("\t\t 3 退出系统")
		fmt.Println("\t\t 请选择（1-3）：")
		fmt.Scanf("%d\n", &key)
		switch key {
		case 1:
			fmt.Println("登陆聊天室")
			fmt.Println("请输入你的ID：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入你的密码：")
			fmt.Scanf("%s\n", &userPwd)
			//init instance
			up := &process.UserProcess{}
			up.Login(userId, userPwd)
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户ID：")
			fmt.Scanf("%d\n", &userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanf("%s\n", &userPwd)
			fmt.Println("请输入用户名：")
			fmt.Scanf("%s\n", &userName)
			up := &process.UserProcess{}
			up.Register(userId, userPwd, userName)
		case 3:
			fmt.Println("退出系统")
			// loop = false
			os.Exit(0)
		default:
			fmt.Println("输入有误，请重新输入")
		}
	}
}