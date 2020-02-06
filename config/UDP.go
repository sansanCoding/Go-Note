package config

//UDP相关配置
type udp struct {
	//UDP监听协议
	NetWork string
	//UDP监听ip
	IP string
	//UDP监听端口
	Port int
}

var UDP *udp

//初始化
func init(){
	UDP = &udp{
		NetWork : "udp",
		IP 		: "127.0.0.1",
		Port	: 30000,
	}
}