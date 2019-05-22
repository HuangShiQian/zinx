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
	//运行服务器
	Serve()
}
