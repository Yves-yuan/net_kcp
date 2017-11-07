package bytebuffer

import (
	"fmt"
	"testing"
	"bytes"
)

func TestByteBuffer(t *testing.T) {


	byteBuf :=  NewWrappedByteBuffer()

	byteBuf.PutUint32(1111)

	fmt.Println(byteBuf.String())
	value,err := byteBuf.GetUint32()
	fmt.Println(value,err)
	byteBuf.PutUint32(1111)
	byteBuf.PutUint32(2222)
	value,err = byteBuf.GetUint32()
	fmt.Println(value,err)
	value,err = byteBuf.GetUint32()
	fmt.Println(value,err)
	value,err = byteBuf.GetUint32()
	fmt.Println(value,err)

	byteBuf.PutUint32(1111)
	byteBuf.PutUint32(2222)

	dst := make([]byte,4)
	err = byteBuf.GetBytes(dst,0,8)


}

func BenchmarkByteBuffer(b *testing.B) {
	byteBuf := NewWrappedByteBuffer()
	b.ReportAllocs()
	for i := 0;i<b.N;i++ {
		byteBuf.PutUint32(1111)
		byteBuf.GetUint32()
	}
}

func BenchmarkBo(b *testing.B) {
	byteBuf := NewWrappedByteBuffer()
	b.ReportAllocs()
	data := make([]byte,4)

	for i := 0;i<b.N;i++ {
		byteBuf.Order().PutUint32(data,1)
		byteBuf.Order().Uint32(data)
	}
}

func BenchmarkByteBuffer1(b *testing.B) {
	byteBuf := NewWrappedByteBuffer()
	data := make([]byte,4)
	byteBuf.Order().PutUint32(data,1)
	byteBuf.PutBytes(bytes.Repeat([]byte{'a','b','c','d'},b.N),0,4*b.N)
	for i := 0;i<b.N;i++ {
		byteBuf.GetUint32()
	}
}