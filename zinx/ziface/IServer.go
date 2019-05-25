/**
 Server模块的抽象层接⼝类
ziface 主要是存放⼀些Zinx框架的全部模块的抽象层接⼝类，Zinx框架的最基本的是服务类接⼝
*/
package ziface

type IServer interface {
	//启动服务器
	Start()
	//停止服务器
	Stop()
	//运行服务器  开启业务服务方法
	Serve()

	//添加路由方法  暴露给开发者的 讓用戶添加router
	//AddRouter(router IRouter)
	AddRouter(msgId uint32,router IRouter)
}
