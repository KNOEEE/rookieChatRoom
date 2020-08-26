package model

import (
	"net"
	"go_code/projectLearn/chatRoom/common/message"
)

type CurUser struct {
	Conn net.Conn
	message.User
}