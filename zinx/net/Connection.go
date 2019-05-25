package net

import (
	"net"
	"zinx/ziface"
	"fmt"
	//"zinx/config"
	"io"
	"errors"
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

	/*//当前链接所绑定的Router
	Router ziface.IRouter*/

	//消息管理模块 多路由
	MsgHandler ziface.IMsgHandler
}

/*
初始化链接方法
 */
func NewConnection(conn *net.TCPConn,connID uint32,handler ziface.IMsgHandler) ziface.IConnection {
	c:=&Connection{
		Conn:conn,
		ConnID:connID,
		//handleAPI:callback_api,
		MsgHandler:handler,
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
		/*
		buf:=make([]byte,config.GlobalObject.MaxPackageSize)
		cnt,err:=c.Conn.Read(buf)
		if err!=nil{
			fmt.Println("recv buf err",err)
			continue
		}
		*/

		//创建拆包封包的对象
		dp:=NewDataPack()

		//读取客户端消息的头部
		headData:=make([]byte,dp.GetHeadLen())////获取二进制包的头部长度  固定返回8
		if _,err:=io.ReadFull(c.Conn,headData);err!=nil{
			fmt.Println("read msg head error",err)
			break
		}

		//根据头部 获取数据的长度，进行第二次读取
		msg,err:=dp.UnPack(headData) //将msg 头部信息填充满
		if err!=nil {
			fmt.Println("unpack error ",err)
			break
		}

		//根据长度 再次读取
		var data []byte
		if msg.GetMsgLen()>0{
			//说明有内容
			data=make([]byte,msg.GetMsgLen())
			if _,err:=io.ReadFull(c.Conn,data);err!=nil{
				fmt.Println("read msg data error  ",err)
				break
			}
		}
		msg.SetMsgData(data)

		//将读出来的msg 组装一个request
		//将当前一次性得到的对端客户端请求的数据 封装成一个Request
		req:=NewRequest(c,msg)
		go c.MsgHandler.DoMsgHandler(req)

		//调用用户传递进来的业务 模板 设计模式
		/*go func() {
			c.Router.PreHandle(req)
			c.Router.Handle(req)
			c.Router.PostHandle(req)
		}()*/

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
func (c *Connection)Send(msgId uint32,msgData []byte) error  {

	if c.isClosed==true{
		return errors.New("Connection closed ..send Msg ")

	}
	//封装成msg
	dp:=NewDataPack()
	binaryMsg,err:=dp.Pack(NewMsgPackage(msgId,msgData))
	if err!=nil{
		fmt.Println("Pack error msg id = ",msgId)
		return err
	}

	//将binaryMsg发送给对端
	if _,err:=c.Conn.Write(binaryMsg);err!=nil{
		fmt.Println("send buf error")
		return err
	}
	return nil
}