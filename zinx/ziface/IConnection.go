package ziface

import "net"

/*
	抽象链接层
 */
type IConnection interface {
	//启动链接
	Start()
	//停止鏈接
	Stop()
	//获取链接ID
	GetConnID() uint32
	//获取conn的原生socket套接字
	GetTcpConnection() *net.TCPConn
	//获取远程客户端的ip地址
	GetRemoteAddr() net.Addr
	//发送数据给对方客户端
	Send([]byte,int) error
}

//业务处理方法 抽象定义   定义⼀个统⼀处理链接业务的接⼝
//HandFunc这个函数类型 是所有conn链接在处理业务的函数接⼝，
/*
第⼀参数是socket原⽣链接，第⼆个参数是客户端请求的
数据，第三个参数是客户端请求的数据⻓度。这样，如果我们想要指定⼀个conn的处理业务，只要定义
⼀个HandFunc类型的函数，然后和该链接绑定就可以了。
 */
//type HandleFunc func(*net.TCPConn,[]byte,int) error
type HandleFunc func(req IRequest) error