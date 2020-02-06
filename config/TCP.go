package config

//TCP相关配置
type tcp struct {
	//TCP监听协议
	NetWork string
	//TCP监听地址
	Address string
}

var TCP *tcp

//初始化
func init(){
	TCP = &tcp{
		NetWork	: "tcp",
		Address	: "127.0.0.1:20000",
	}
}