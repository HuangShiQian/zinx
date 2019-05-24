package net

import (
	"net"
	"zinx/ziface"
	"fmt"
	"zinx/config"
)

//具体的TCP链接模块
type Connection struct {
	//当前链接的原生套接字  TCP套接字
	Conn *net.TCPConn

	//当前连接的ID 也可以称作为SessionID，ID全局唯⼀
	ConnID uint32

	//当前的链接状态
	isClosed bool

	//当前链接所绑定的业务处理方法
	//handleAPI ziface.HandleFunc

	//当前链接所绑定的Router
	Router ziface.IRouter
}

/*
初始化链接方法
 */
func NewConnection(conn *net.TCPConn,connID uint32,router ziface.IRouter) ziface.IConnection {
	c:=&Connection{
		Conn:conn,
		ConnID:connID,
		//handleAPI:callback_api,
		Router:router,
		isClosed:false,
	}
	return c
}

//针对链接读业务的方法  处理conn读数据的Goroutine
func (c *Connection)StartReader()  {
	//从对端读数据
	fmt.Println("Reader go is startin....")

	defer fmt.Println("connID = ",c.ConnID,"Reader is exit, remote addr is =",c.GetRemoteAddr().String())
	defer c.Stop()

	for  {
		buf:=make([]byte,config.GlobalObject.MaxPackageSize)
		cnt,err:=c.Conn.Read(buf)
		if err!=nil{
			fmt.Println("recv buf err",err)
			continue
		}
		//将当前一次性得到的对端客户端请求的数据 封装成一个Request
		req:=NewRequest(c,buf,cnt)

		//调用用户传递进来的业务 模板 设计模式
		go func() {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}()

		/*
		//将数据 传递给我们 定义好的Handle Callback方法
		if err:=c.handleAPI(req);err!=nil{
			fmt.Println("ConnID",c.ConnID,"Handle is error",err)
			break
		}
		*/
	}
}

//启动链接
func (c *Connection)Start()  {
	fmt.Println("Conn Start（）  ... id =",c.ConnID)
	//先进行读业务
	go c.StartReader()
	//TODO 进行写业务

}

//停止链接
func (c *Connection)Stop()  {
	fmt.Println("c. Stop() ... ConnId = ",c.ConnID)
	//回收工作
	if c.isClosed==true{
		return
	}
	c.isClosed=true

	//关闭原生套接字
	_=c.Conn.Close()

}

//获取链接ID
func (c *Connection)GetConnID() uint32 {
	return c.ConnID

}

//获取conn的原生socket套接字
func (c *Connection)GetTcpConnection() *net.TCPConn  {
	return c.Conn
}

//获取远程客户端的ip地址
func (c *Connection)GetRemoteAddr() net.Addr  {
	return c.Conn.RemoteAddr()
}

//发送数据给对方客户端
func (c *Connection)Send(data []byte,cnt int) error  {
	if _,err:=c.Conn.Write(data[:cnt]);err!=nil{
		fmt.Println("send buf error")
		return err
	}
	return nil
}