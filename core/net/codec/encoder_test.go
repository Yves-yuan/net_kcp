package codec

import (
	"bytes"
	"server/message"
	"strings"
	"testing"

	"github.com/golang/protobuf/proto"
)

func BenchmarkCTransportEncoder_Encode(b *testing.B) {
	b.ReportAllocs()

	encoder := NewTransportEncoder()
	data := bytes.Repeat([]byte("a"), 100)
	for i := 0; i <= b.N; i++ {
		encoder.Encode(1, 0, data[:i%100])
	}
}

func BenchmarkProtobufSerialize(b *testing.B) {
	b.ReportAllocs()
	msg := &message.ReqAllBooksInfo{
		Token: strings.Repeat("s", 100),
	}
	for i := 0; i <= b.N; i++ {
		if _, err := proto.Marshal(msg); err != nil {
			b.Error(err)
		}
	}
}
