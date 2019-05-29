package main

import (
	"ZinxDemo/protobufDemo/pb"
	"github.com/golang/protobuf/proto"
	"fmt"
)

func main()  {
	//定义一个protobuf结构体对象
	person:=&pb.Person{
		Name:"Rose",
		Age:16,
		Emails:[]string{"17327735377@163.com","15751833577@163.com"},
		Phones:[]*pb.PhoneNumber{
			&pb.PhoneNumber{
				Number:"13666666666",
			},
			&pb.PhoneNumber{
				Number:"13888888888",
			},
			&pb.PhoneNumber{
				Number:"13999999999",
			},
		},

		//oneof赋值
		Data:&pb.Person_School{
			School:"中国人民大学",
		},
	}

	//将一个protobuf结构体对象 转化成二进制数据
	//任何proto message结构体 在go中他们都是基于Message接口的

	//编码
	data,err:=proto.Marshal(person)
	if err!=nil{
		fmt.Println("marshal err ",err)
		return
	}
	//data就是我们要刚给对端发送的二进制数据

	//对端已经收到了data了
	//解码
	newPerson:=&pb.Person{}
	err=proto.Unmarshal(data,newPerson)//将data解码值 newPerson结构体中
	if err!=nil{
		fmt.Println("unmarshal err ",err)
		return
	}
	fmt.Println("源数据：",person)
	fmt.Println("解码之后数据:",newPerson)

	fmt.Println("name = ",newPerson.GetName(),"age = ",newPerson.GetAge(),"emails: ",newPerson.GetEmails(),"numbers = ",newPerson.GetPhones())
	fmt.Println("School = ",newPerson.GetSchool())
}