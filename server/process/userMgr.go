package prcs
import (
	"fmt"
)

//因为Mgr实例在服务器端仅有一个 将其定义为全局变量
var (
	//这里感觉应该加个锁啊 并发的时候会访问冲突吧
	userMgr *UserMgr 
)

//用于显示在线用户list
type UserMgr struct {
	onlineUsers map[int]*UserProcess
}

func init() {
	userMgr = &UserMgr{
		onlineUsers : make(map[int]*UserProcess, 1024),
	}
}

//添加
func(this *UserMgr) AddOnlineUser(up *UserProcess){
	this.onlineUsers[up.UserId] = up
}

func(this *UserMgr) DelOnlineUser(userId int){
	delete(this.onlineUsers, userId)
}

//返回当前all在线用户
func(this *UserMgr) GetAllOnlineUsers() map[int]*UserProcess {
	return this.onlineUsers
}

//根据id返回
func(this *UserMgr) GetOnlineUserById(userId int) (up *UserProcess,
	err error) {
	up, ok := this.onlineUsers[userId]
	if !ok {
		//要查找的用户当前不在线
		err = fmt.Errorf("用户%d离线", userId)
		return
	}
	return
}