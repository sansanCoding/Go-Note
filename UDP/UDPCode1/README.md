# UDPCode1
第一版简易版UDP测试代码,联调下服务端与客户端通信处理.

```sh
# 调用顺序:
# 1.先执行UDP服务端的开启监听
#   运行示例:
$   go run ./main.go
#   然后再输入如下字符开启TCP服务端监听
$   {"optTag":"UDPCode1","optParams":{"doTag":"serverStart"}}
# 2.客户端发起消息
#   由于上面第1步已经进入到服务端监听,如下输入客户端消息发送的字符,其中sendMsg就是发送的具体消息
$   {"optTag":"UDPCode1","optParams":{"doTag":"clientStart","sendMsg":"Test"}}
```