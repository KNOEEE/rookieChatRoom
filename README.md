# rookieChatRoom
go实现聊天室入门

需求：
1.用户注册
2.用户登录
3.显示在线用户列表
4.群聊（广播）

测试时，先在redis里加一个用户
hset users 100 "{\"userId\":100,\"userPwd\":\"abc\",\"userName\":\"Scott\"}"
