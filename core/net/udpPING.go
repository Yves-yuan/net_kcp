package net

import (
	"encoding/binary"
	"net"
	"time"
)

const UDPdebug = false

func StartudpPing(port int) {
	// 开启端口监听
	socket, err := net.ListenUDP("udp4", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: port,
	})

	if err != nil {
		logger.Error("UDP端口监听失败 :", err)
		return
	}
	defer socket.Close()

	logger.Info("UDP ping 端口监听:", port)
	loop(socket)

}

func loop(socket *net.UDPConn) bool {
	// ping包大小不应该大于MTU
	data := make([]byte, 576)
	for {
		// 接受数据
		number, remoteAddr, err := socket.ReadFromUDP(data)
		if err != nil {
			logger.Error("读取数据失败 :", err)
			continue
		}

		temp := make([]byte, number)
		copy(temp, data)

		time := time.Now().UnixNano()
		timebyte := make([]byte, 8)
		binary.BigEndian.PutUint64(timebyte, uint64(time))

		sendData := append(temp, timebyte...)
		if UDPdebug == true {
			logger.Infof("来自%s 的ping请求,字节=%d 服务器时间:%d, data:%b", remoteAddr, number, time, sendData)
		}

		// 收到数据后，立刻发送
		_, err = socket.WriteToUDP(sendData, remoteAddr)
		if err != nil {
			logger.Error("发送数据失败 :", err)
			continue
		}
	}

	return true
}
