/**
server模块的实现层
net 模块是zinx框架中⽹络相关功能的实现，所有⽹络相关模块都会定义在 znet 模块中
*/
package net

import (
	"zinx/ziface"
	"fmt"
	"net"

	"zinx/config"
)

type Server struct {
	//服务器ip
	IPVersion string
	IP string
	//服务器port
	Port int
	//服务器名称
	Name string
	/*//路由属性
	Router ziface.IRouter*/

	//多路由的消息管理模块
	MsgHandler ziface.IMsgHandler

	//链接管理模块
	connMgr ziface.IConnManager

	//该server创建链接之后自动调用Hook函数
	OnConnStart func(conn ziface.IConnection)
	//该server销毁链接之前自动调用的Hook函数
	OnConnStop func(conn ziface.IConnection)
}

/*//定义一个 具体的回显业务 针对type HandleFunc func(*net.TCPConn,[]byte,int) error
func CallBackBusi(request ziface.IRequest)error  {
	//回显业务
	fmt.Println("【conn Handle】 CallBack..")
	c:=request.GetConnection().GetTcpConnection()
	buf:=request.GetData()
	cnt:=request.GetDataLen()
	if _,err:=c.Write(buf[:cnt]);err!=nil{
		fmt.Println("write back err ",err)

		return err
	}
	return nil
}*/

//初始化的New方法
func NewServer(name string)ziface.IServer  {
	s:=&Server{
		Name:config.GlobalObject.Name,
		IPVersion:"tcp4",
		IP:config.GlobalObject.Host,
		Port:config.GlobalObject.Port,
		MsgHandler:NewMsgHandler(),
		connMgr:NewConnManager(),
	}

	return s
}

//启动服务器
//原生socket 服务器编程
func (s *Server)Start()  {
fmt.Printf("[start] Server Linstenner at IP :%s ,Port:%d ,is starting...\n",s.IP,s.Port)

	//0 启动worker工作池
	s.MsgHandler.StartWorkerPool()

	//1 创建套接字  ：得到一个TCP的addr
	/*
	ResolveTCPAddr将addr作为TCP地址解析并返回。
	参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；net必须是"tcp"、"tcp4"或"tcp6"。
	 */
	addr,err:=net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))
	if err!=nil{
		fmt.Println("resolve tcp addr error:",err)
		return
	}

	//2 监听服务器地址
	/*
	ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener，net参数必须是"tcp"、"tcp4"、"tcp6"，
	如果laddr的端口字段为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。
	 */
	listenner, err:=net.ListenTCP(s.IPVersion,addr)
	if err!=nil{
		fmt.Println("listen",s.IPVersion,"err:",err)
		return
	}

	//生成id的累加器
	var cid uint32
	cid = 0

	//3 阻塞等待客户端发送请求，
	go func() {
		for  {
			//阻塞等待客户端请求,  AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn。
			conn,err:=listenner.AcceptTCP()//只是针对TCP协议
			if err!=nil {
				fmt.Println("Accept err",err)
				continue
			}

			//创建一个Connection对象
			//判断当前server链接数量是否已经最大值
			if s.connMgr.Len() >= int(config.GlobalObject.MaxConn){
				//当前链接已经满了
				fmt.Println("---> Too Many Connection, MaxConn = ",config.GlobalObject.MaxConn)
				conn.Close()
				continue
			}
			dealConn:=NewConnection(s,conn,cid,s.MsgHandler)//Router和連接建立联系
			cid++

			//此时conn就已经和对端客户端连接
			go dealConn.Start()

			/*//此时conn就已经和对端客户端连接
			go func() {
				//4 客户端有数据请求，处理客户端业务(读、写)
				for  {
					buf:=make([]byte,512)
					cnt,err:=conn.Read(buf)
					if err!=nil&&err!=io.EOF{
						fmt.Println("recv buff err",err)//EOF
						break
					}
					fmt.Printf("recv client buf %s,cnt=%d\n",buf[:cnt],cnt)

					//回显功能 （业务）
					if _,err=conn.Write(buf[:cnt]);err!=nil{
						fmt.Println("write back buf err",err)
						continue
					}
				}
			}()*/
		}
	}()
}

//停止服务器
func (s *Server)Stop()  {
	//服务器停止  应该清空当前全部的链接
	s.connMgr.ClearConn()
}

//运行服务器
func (s *Server)Serve()  {
	//启动server的监听功能
	s.Start()//并不希望他永久的阻塞

	//TODO  做一些其他的扩展
	//阻塞//告诉CPU不再需要处理的，节省cpu资源
	select {} //main函数不退出

}

//添加路由方法  暴露给开发者的 [IServer裏面有 來這邊要實現]
func (s *Server) AddRouter(msgId uint32,router ziface.IRouter) {

	//s.Router = router   //左邊router是PingRouter  它實現了三個方法
	s.MsgHandler.AddRouter(msgId,router)
	fmt.Println("Add Router SUCC!! msgID = ",msgId)
}

//提供一个得到链接管理模块的方法
func (s *Server)GetConnMgr() ziface.IConnManager{
	return s.connMgr
}

//注册 创建链接之后 调用的 Hook函数 的方法
func (s *Server)AddOnConnStart(hookFunc func(coon ziface.IConnection)){
	s.OnConnStart=hookFunc
}
//注册 销毁链接之前调用的Hook函数 的方法
func (s *Server)AddOnConnStop(hookFunc func(coon ziface.IConnection)){
	s.OnConnStop=hookFunc
}
//调用 创建链接之后的HOOK函数的方法
func (s *Server)CallOnConnStart(conn ziface.IConnection){
	if s.OnConnStart!=nil{
		fmt.Println("---> Call OnConnStart()...")
		s.OnConnStart(conn)
	}
}
//调用 销毁链接之前调用的HOOk函数的方法
func (s *Server)CallOnConnStop(conn ziface.IConnection){
	if s.OnConnStop!=nil{
		fmt.Println("---> Call OnConnStop()...")
		s.OnConnStop(conn)
	}
}