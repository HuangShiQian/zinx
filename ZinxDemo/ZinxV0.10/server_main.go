package main

import (
	"zinx/net"
	"zinx/ziface"
	"fmt"
)
//PreHandle方法  ---  用户可以在处理业务之前  自定义一些业务， 实现这个方法
//Handler方法  ---- 用户可以定义一个 业务处理的 核心方法
//PostHandle方法  --- 用户可以在处理业务之后 定义一些业务，实现这个方法
type PingRouter struct {
	net.BaseRouter
}


/*

//提供自定义的业务方法
func (this *PingRouter)PreHandle(request ziface.IRequest)  {
	fmt.Println("Call Router PreHandler...")
	//给客户端回写一个 数据
	_,err:=request.GetConnection().GetTcpConnection().Write([]byte("before ping...\n"))
	if err!=nil{
		fmt.Println("call back before ping error")
	}
}
*/


//200 ---> pingpingping
//201 ---> hello zinx..


func(this *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handler...")
	//给客户端回写一个 数据
	err:=request.GetConnection().Send(200,[]byte("ping..ping..ping...\n"))
	if err!=nil{
		fmt.Println(err)
	}
}
/*

func (this *PingRouter)PostHandle(request ziface.IRequest)  {
	fmt.Println("Call Router PostHandler...")
	//给客户端回写一个 数据
	_,err:=request.GetConnection().GetTcpConnection().Write([]byte("after ping...\n"))
	if err!=nil{
		fmt.Println("call back after ping error")
	}
}
*/

type HelloRouter struct {
	net.BaseRouter
}


func (this *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call Router Handler...")
	//给客户端回写一个 数据
	err:=request.GetConnection().Send(201,[]byte("Hello Zinx!!!"))
	if err!=nil{
		fmt.Println(err)
	}
}

//创建链接之后的执行的钩子函数
func DoConnectionBegin(conn ziface.IConnection)  {
	fmt.Println("===> DoConnectionBegin  ....")
	//链接一旦创建成功 给用户返回一个消息
	if err:=conn.Send(202,[]byte("Hello Welcome to zinx..."));err!=nil{
		fmt.Println(err)
	}

	//当用户一旦链接创建成功， 给链接绑定一些属性
	fmt.Println("Set conn property...")
	conn.SetProperty("Name","Go3")
	conn.SetProperty("Address","TBD")
	conn.SetProperty("Time","2019-06-06")
}

//链接销毁之前执行的钩子函数
func DoConnectionLost(conn ziface.IConnection)  {
	fmt.Println("===> DoConnectionLost  ....")
	//链接一旦断开 给用户返回一个消息
	fmt.Println("Conn id ",conn.GetConnID(),"is Lost!...")

	fmt.Println("Get Conn Property...")
	//获取conn Name
	if name,err:=conn.GetProperty("Name");err==nil{
		fmt.Println("Name =",name)
	}
	//获取conn address
	if address,err:=conn.GetProperty("Address");err==nil{
		fmt.Println("Address =",address)
	}
	//获取conn time
	if time,err:=conn.GetProperty("Time");err==nil{
		fmt.Println("Time =",time)
	}
}


func main()  {
	//创建一个zinx server对象
	s:=net.NewServer("zinx v0.10")

	//注册一个创建链接之后的方法业务
	s.AddOnConnStart(DoConnectionBegin)
	//注册一个链接断开之前的方法业务
	s.AddOnConnStop(DoConnectionLost)

	//注册一些自定义的业务
	//s.AddRouter(&PingRouter{})
	s.AddRouter(1,&PingRouter{})
	s.AddRouter(2,&HelloRouter{})

	//让server对象 启动服务
	s.Serve()

	return
}
