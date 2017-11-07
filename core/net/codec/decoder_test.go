package codec

import (
	"strings"
	"testing"

	"encoding/binary"
	"server/message"
	"server/route"
	"time"

	"github.com/golang/protobuf/proto"
)

func TestCTransportDecoder_Decode(t *testing.T) {
	route.Register(message.MSGID_ReqTestMessageE, test, (*message.ReqAllBooksInfo)(nil))

	token := strings.Repeat("1", testMsgLength)
	data, err := proto.Marshal(&message.ReqAllBooksInfo{Token: token})
	if err != nil {
		t.Error(err)
	}

	count := 2
	buffer := make([]byte, (len(data)+HEADER_SIZE)*count)
	total := 0

	for i := 0; i < count; i++ {
		mlen := len(data) + HEADER_SIZE - recvSizeLength
		binary.BigEndian.PutUint32(buffer[total:total+4], uint32(mlen))
		binary.BigEndian.PutUint64(buffer[total+4:total+12], uint64(time.Now().Nanosecond()))

		order := i ^ (0xFE98 << 8)
		order = order ^ mlen
		binary.BigEndian.PutUint32(buffer[total+12:total+16], uint32(order))
		binary.BigEndian.PutUint16(buffer[total+16:total+18], uint16(message.MSGID_ReqTestMessageE))

		//println(len(buffer), total+HEADER_SIZE, total+HEADER_SIZE+len(data), total, len(data))
		copy(buffer[total+HEADER_SIZE:total+HEADER_SIZE+len(data)], data)
		total += mlen + 4
	}

	decoder := NewTransportDecoder()
	out := decoder.Decode(nil, buffer)
	if len(out) < 1 {
		t.Fail()
	}

	for _, item := range out {
		got := item.Payload().(*message.ReqAllBooksInfo).Token
		if token != got {
			t.Errorf("expect: %s, got: %s", token, got)
		}
	}
	route.UnRegister(message.MSGID_ReqTestMessageE)
}

func BenchmarkCTransportDecoder_Decode(b *testing.B) {
	route.Register(message.MSGID_ReqTestMessageE, test, (*message.ReqAllBooksInfo)(nil))

	token := strings.Repeat("1", testMsgLength)
	data, err := proto.Marshal(&message.ReqAllBooksInfo{Token: token})
	if err != nil {
		b.Error(err)
	}
	buffer := make([]byte, len(data)+HEADER_SIZE)

	mlen := len(data) + HEADER_SIZE - recvSizeLength
	binary.BigEndian.PutUint32(buffer[0:4], uint32(mlen))
	binary.BigEndian.PutUint64(buffer[4:12], uint64(time.Now().Nanosecond()))

	order := 0 ^ (0xFE98 << 8)
	order = order ^ mlen
	binary.BigEndian.PutUint32(buffer[12:16], uint32(order))
	binary.BigEndian.PutUint16(buffer[16:18], uint16(message.MSGID_ReqTestMessageE))

	//println(len(buffer), total+HEADER_SIZE, total+HEADER_SIZE+len(data), total, len(data))
	copy(buffer[HEADER_SIZE:HEADER_SIZE+len(data)], data)

	b.ReportAllocs()

	decoder := NewTransportDecoder()
	for i := 0; i < b.N; i++ {
		out := decoder.Decode(nil, buffer)
		if len(out) < 1 {
			b.Fatal("decode error")
		}

		got := out[0].Payload().(*message.ReqAllBooksInfo).Token
		if token != got {
			b.Fatalf("expect: %s, got: %s", token, got)
		}
		decoder.Reset()
	}

	route.UnRegister(message.MSGID_ReqTestMessageE)
}

func BenchmarkHandlerPayload(b *testing.B) {
	route.Register(message.MSGID_ReqTestMessageE, test, (*message.ReqAllBooksInfo)(nil))

	token := strings.Repeat("1", testMsgLength)
	data, err := proto.Marshal(&message.ReqAllBooksInfo{Token: token})
	if err != nil {
		b.Error(err)
	}

	b.ReportAllocs()
	handler, err := route.FetchHandler(message.MSGID_ReqTestMessageE)
	if err != nil {
		b.Fail()
	}

	for i := 0; i < b.N; i++ {
		payload := handler.Payload()
		if err := proto.Unmarshal(data, payload); err != nil {
			b.Fatal(err)
		}
	}

	route.UnRegister(message.MSGID_ReqTestMessageE)
}
