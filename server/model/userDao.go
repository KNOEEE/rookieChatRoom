package model
import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go_code/projectLearn/chatRoom/common/message"
	"encoding/json"
)

//在服务器启动后 就初始化一个dao实例 做成一个全局变量
var (
	MyUserDao *UserDao
)

//定义userDao结构体 完成对User结构体的操作
type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式 创建一个userdao的实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao){
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//根据用户id返回用户实例+err
func(this *UserDao) getUserById(conn redis.Conn, 
	id int) (user *User, err error) {
	//通过给定id去redis查询
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			//表示在usershash中没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &User{}
	//把res反序列化成user实例
	json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err =", err)
		return
	}
	return
}

//完成登陆的校验
//如果id和pwd都正确 则返回一个user实例
//如果不正确 则返回对应的错误信息
func(this *UserDao) Login(userId int, 
	userPwd string) (user *User, err error) {
	//从pool中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//获取到用户 开始校验密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}

func(this *UserDao) Register(user *message.User) (err error) {
	//从pool中取出一个连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	//说明该用户还未注册过 则可以注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}
	_, err = conn.Do("HSet", "users", 
		user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户 err =", err)
		return
	}
	return
}