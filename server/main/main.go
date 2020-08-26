package main
import (
	"fmt"
	"net"
	"time"
	"go_code/projectLearn/chatRoom/server/model"
)
//go build -o server.exe go_code\projectLearn\chatRoom\server\main

//处理和客户端的通信
func process(conn net.Conn) {
	defer conn.Close() // 及时关闭！！
	processor := &Processor{
		Conn: conn,
	}
	err := processor.processPro()
	if err != nil {
		fmt.Println("server通信协程err =", err)
		return
	}
}

func init() {
	//服务器启动时 先初始化redis连接池子
	//研究了半天逻辑 两边通讯都会卡住 原来是redis一开始就没连接上
	//redis看起来连着 好像隔了太长时间 自己断开了
	initPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao()
}

//编写一个函数 完成对userDao的初始化
func initUserDao() {
	//pool本身就是全局变量
	//该函数必须在initPool后面
	model.MyUserDao = model.NewUserDao(pool)
}

func main() {
	fmt.Println("server is listening port 8888...")
	listen, err := net.Listen("tcp", "localhost:8888")
	defer listen.Close()

	if err != nil {
		fmt.Println("net.Listen err =", err)
		return
	}

	for {
		fmt.Println("Waiting cli connecting...")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("listen.Accept err =", err)
		}

		//启动一个协程和cli保持通信
		go process(conn)
	}
}