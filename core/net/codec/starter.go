package codec

import (
	"server/core/message/packet"
	"server/core/net/internal/kcp"

	"server/core/log"
)

var logger = log.New("codec")

func StartDecode(in <-chan *kcp.SessionPacket, out chan<- *packet.CRecvPacket) {
	for {
		select {
		case p, ok := <-in:
			if !ok {
				return
			}

			list := p.From.Decoder().Decode(p.From.Session(), p.Data)
			for _, item := range list {
				out <- item
			}
		}
	}
}

func StartEncode(in <-chan *packet.CSendPacket, out chan<- *kcp.SessionPacket) {
	for {
		select {
		case p, ok := <-in:
			if !ok {
				return
			}

			payload, err := p.Payload()
			if err != nil {
				logger.Error(err)
				continue
			}

			udpSession := p.Conn().(*kcp.UDPSession)
			data := udpSession.Encoder().Encode(p.ID(), p.Type(), payload)

			out <- &kcp.SessionPacket{udpSession, data}
		}
	}
}
