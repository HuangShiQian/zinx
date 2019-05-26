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



func main()  {
	//创建一个zinx server对象
	s:=net.NewServer("zinx v0.8")

	//注册一些自定义的业务
	//s.AddRouter(&PingRouter{})
	s.AddRouter(1,&PingRouter{})
	s.AddRouter(2,&HelloRouter{})

	//让server对象 启动服务
	s.Serve()

	return
}
