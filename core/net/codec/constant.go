package codec

// 接收消息协议头定义
// +----------------------------------------------------------------+
// | SIZE(4) | TIMESTAMP(8) | ORDER(4) | PROTO(2) | PAYLOAD(SIZE-14)|
// +----------------------------------------------------------------+
const (
	recvSizeLength    = 4                                  // 协议中长度字段的字节数
	recvTimeOffset    = 0                                  // 协议头中时间字段偏移位置(不计算长度的那一部分)
	recvTimeSize      = 8                                  // 协议中时间字段的字节数
	recvOrderOffset   = recvTimeOffset + recvTimeSize      // 协议中序号字段的偏移位置
	recvOrderLength   = 4                                  // 协议中序号字段的字节数
	recvProtoOffset   = recvOrderOffset + recvOrderLength  // 协议头中消息ID字段偏移位置
	recvProtoSize     = 2                                  // 协议中消息类型的字节数
	recvPayloadOffset = recvProtoOffset + recvProtoSize    // payload偏移位置
	HEADER_SIZE       = recvPayloadOffset + recvSizeLength // 协议头总长度
	MAX_RECV_SIZE     = 20 * 1024                          // 最大接收字节数
	MAX_SEND_SIZE     = 20 * 1024                          // 最大发送字节数
)

// 发送消息协议头定义
// +----------------------------------------------------------------+
// | SIZE(4) | TYPE(1) | PROTO(2)  | PAYLOAD(SIZE-3)                |
// +----------------------------------------------------------------+
const (
	sendSizeLength    = 4                                 // 消息长度字段字节数
	sendTypeOffset    = sendSizeLength                    // 消息类型偏移位置
	sendTypeLength    = 1                                 // 消息类型字节数
	sendProtoOffset   = sendTypeOffset + sendTypeLength   // 消息ID偏移位置
	sendProtoLength   = 2                                 // 消息ID字段字节数
	sendPayloadOffset = sendProtoOffset + sendProtoLength // Payload偏移位置
)
