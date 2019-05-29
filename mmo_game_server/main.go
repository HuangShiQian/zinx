package main

import (
	"zinx/net"
	"zinx/ziface"
	"fmt"
	"mmo_game_server/core"
)

//当前客户端建立链接之后触发Hook函数
func OnConnectionAdd(conn ziface.IConnection)  {
	fmt.Println("Conn Add...")

	//创建一个玩家 将链接和玩家模块绑定
	p:=core.NewPlayer(conn)

	//给客户端发送一个msgID:1
	p.ReturnPid()

	//给客户端发送一个msgID:200
	p.ReturnPlayerPosition()

	//上线成功了
	//将玩家对象添加到世界管理器中
	core.WorldMngrObj.AddPlayer(p)

	fmt.Println("----> Player ID = ",p.Pid,"Online...",", Player num = ",len(core.WorldMngrObj.Players))
}

func main()  {
	s:=net.NewServer("MMO Game Server")

	//注册一些 链接创建/销毁的 Hook钩子函数
	s.AddOnConnStart(OnConnectionAdd)

	//注册一些路由业务

	s.Serve()
}
