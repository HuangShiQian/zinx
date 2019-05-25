package main

import (
	"fmt"
	"time"
	"net"
	net2 "zinx/net"
	"io"
)

/*
	模拟客户端
 */
func main()  {
	fmt.Println("client start...")

	time.Sleep(1*time.Second)
	//直接connect 服务器得到一个 已经建立好的conn句柄
	conn,err:=net.Dial("tcp","127.0.0.1:8999")
	if err !=nil{
		fmt.Println("client start err",err)
		return
	}

	for   {
		/*//写
		_,err=conn.Write([]byte("hello zinx ..."))
		if err!=nil{
			fmt.Println("conn write err",err)
			return
		}

		//读
		buf:=make([]byte,512)
		cnt,err:=conn.Read(buf)
		if err!=nil{
			fmt.Println("conn read err",err)
			return
		}

		fmt.Printf("server call back : %s,cnt=%d\n",buf[:cnt],cnt)*/

		dp:=net2.NewDataPack()
		binaryMsg,err:=dp.Pack(net2.NewMsgPackage(1,[]byte("pong...pong...pong")))
		if err!=nil{
			fmt.Println("Pack error ",err)
			return
		}
		if _,err:=conn.Write(binaryMsg);err!=nil{
			fmt.Println("write error",err)
			return
		}

		//服务器就会给我们返回一个 消息ID 1 的 pingping TLV格式的二进制数据
		binaryHead:=make([]byte,dp.GetHeadLen())
		if _,err:=io.ReadFull(conn,binaryHead);err!=nil{
			fmt.Println("client unpack msgHead error",err)
			return
		}

		//根据头的长度进行第二次读取
		msgHead,err:=dp.UnPack(binaryHead)//msgHead 是一个IMessage 里面有len 和id
		if msgHead.GetMsgLen()>0{
			//读取包体
			msg:=msgHead.(*net2.Message)
			msg.Data=make([]byte,msg.GetMsgLen())
			if _,err:=io.ReadFull(conn,msg.Data);err!=nil{
				fmt.Println("read msg data error",err)
				return
			}
			fmt.Println("---> Recv Server Msg : id = ",msg.Id,"len = ",msg.DataLen," data = ",string(msg.Data))
		}

		time.Sleep(1*time.Second)
	}
}
