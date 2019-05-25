package net

import "zinx/ziface"

type Message struct {
	Id uint32
	DataLen uint32
	Data []byte
}

//提供一个创建Message的方法
func NewMsgPackage(id uint32,data []byte)ziface.IMessage  {
	return &Message{
		Id:id,
		DataLen:uint32(len(data)),
		Data:data,
	}
}

//getter
func (m *Message)GetMsgId() uint32{
	return m.Id
}
func (m *Message)GetMsgLen() uint32{
	return m.DataLen
}
func (m *Message)GetMsgData() []byte{
	return m.Data
}

//setter
func (m *Message)SetMsgId(id uint32){
	m.Id=id
}
func (m *Message)SetMsgData(data []byte){
	m.Data=data
}
func (m *Message)SetDataLen(len uint32){
	m.DataLen=len
}