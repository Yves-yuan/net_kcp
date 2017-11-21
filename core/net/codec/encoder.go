/*************************************************
  Copyright (C), 2006-2010, 8760 Tech. Co., Ltd.
  作    者: Yuan
  版    本: v1
  完成日期: 2017-07-08
  功能描述: 将序列化好的数据进行封包, 然后返回
*************************************************/
package codec

import (
	"encoding/binary"
)

type CTransportEncoder struct{}

func NewTransportEncoder() *CTransportEncoder {
	return &CTransportEncoder{}
}

// 给数据添加协议头, 包含消息ID, 类型
// msgid:   消息ID
// typ:     消息类型
// payload: 数据
func (c *CTransportEncoder) Encode(msgid, typ int, payload []byte) []byte {
	size := sendProtoLength + sendTypeLength + len(payload)

	data := make([]byte, size+sendSizeLength)

	// 消息长度
	binary.BigEndian.PutUint32(data[:sendSizeLength], uint32(size))

	// 消息类型
	data[sendTypeOffset] = byte(typ)

	// 消息ID
	binary.BigEndian.PutUint16(data[sendProtoOffset:sendProtoOffset+sendProtoLength], uint16(msgid))
	copy(data[sendPayloadOffset:], payload)

	return data
}
