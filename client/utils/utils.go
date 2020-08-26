package utils
import (
	"fmt"
	"net"
	"go_code/projectLearn/chatRoom/common/message"
	"encoding/binary"
	"encoding/json"
)

//将这些方法关联到结构体中
type Transfer struct {
	Conn net.Conn
	//传输时使用的缓冲
	Buf [8096]byte // 数组当切片用
}

func (this *Transfer) ReadPkg() (msg message.Message,
	err error) {
	// buf := make([]byte, 8096)
	// fmt.Println("等待客户端信息")
	//在conn没有关闭时 read 会阻塞
	//任一方关闭连接 则不会阻塞
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	//将buf[:4]转成uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkgLen读取消息内容
	//从conn中读pkgLen个字节到buf中
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}

	//把buf反序列化 注意&符号
	//技术就是一层窗户纸
	json.Unmarshal(this.Buf[:pkgLen], &msg)
	if err != nil {
		fmt.Println("json.Unmarshal fail", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	//先发送一个长度给对方
	var pkgLen uint32
	pkgLen = uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(len) err =", err)
		return
	}

	//发送data本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(data) err =", err)
		return
	}
	return
}