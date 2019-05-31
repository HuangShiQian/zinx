package apis

import (
	"zinx/net"
	"zinx/ziface"
	"mmo_game_server/pb"
	"github.com/golang/protobuf/proto"
	"fmt"
	"mmo_game_server/core"
)

//业务更新坐标 路由业务
type Move struct {
	net.BaseRouter
}

func (m *Move)Handle(request ziface.IRequest)  {
	//解析客户端发送过来的proto协议 msgID:3
	proto_msg:=&pb.Position{}
	proto.Unmarshal(request.GetMsg().GetMsgData(),proto_msg)

	//通过链接属性 得到当前玩家的ID
	pid,_:=request.GetConnection().GetProperty("pid")
	fmt.Println("player id = ",pid.(int32)," move --> ",proto_msg.X,", ",proto_msg.Z,", ",proto_msg.V)

	//通过pid 得到当前的玩家对象
	player:=core.WorldMngrObj.GetPlayerById(pid.(int32))
	//玩家对象方法(将当前的新坐标位置 发送给全部的周边玩家)
	player.UpdatePosition(proto_msg.X,proto_msg.Y,proto_msg.Z,proto_msg.V)
}