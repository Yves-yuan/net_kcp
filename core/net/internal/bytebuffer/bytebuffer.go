/*
 * Copyright (c) 2013 Zhen, LLC. http://zhen.io. All rights reserved.
 * Use of this source code is governed by the Apache 2.0 license.
 *
 */

// ByteBuffer is a Go implementation of the Java ByteBuffer
// http://docs.oracle.com/javase/7/docs/api/java/nio/ByteBuffer.html#put(int, byte)
//
// The documentation for each of the methods are copied from the above page as reference.

// A byte buffer.
// This class defines six categories of operations upon byte buffers:
//
// - Absolute and relative get and put methods that read and write single bytes;
// - Relative bulk get methods that transfer contiguous sequences of bytes from this buffer into an array;
// - Relative bulk put methods that transfer contiguous sequences of bytes from a byte array or some other
//   byte buffer into this buffer;
// - Absolute and relative get and put methods that read and write values of other primitive types, translating
//   them to and from sequences of bytes in a particular byte order;
// - Methods for creating view buffers, which allow a byte buffer to be viewed as a buffer containing values of
//   some other primitive type; and
// - Methods for compacting, duplicating, and slicing a byte buffer.
//
// Byte buffers can be created either by allocation, which allocates space for the buffer's content, or by
// wrapping an existing byte array into a buffer.
//
// Buffers are not safe for use by multiple concurrent threads. If a buffer is to be used by more than one
// thread then access to the buffer should be controlled by appropriate synchronization.
//
// Differences
// -----------
// This is not a full implementation of the Java ByteBuffer so please review the methods yourself.
// Here are a few immediate differences:
// - There is no direct vs indirect buffers in this implementation. All buffers are backed by []byte.
// - Most of the functions cannot be chained like the Java version
// - Some of the functions are not implemented, partly I don't have any use, partly coz I am lazy
//   - array(), arrayOffset()
//   - getChar(), putChar(), asCharBuffer()
//   - Any of the float or double functions
//   - get(byte[] dst), put(byte[] dst)
//   - isDirect()
//   - compact()
//   - hashCode()
//   - compareTo()
//   - view buffers are not implemented except for readonly buffer
// - Some of the functions are renamed to their Go equivalent
//   - Any Short became Uint16, e.g., getShort() -> GetUint16()
//   - Any Int became Uint32
//   - Any Long became Uint64

package bytebuffer

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

const ()

// A buffer is a linear, finite sequence of elements of a specific primitive type. The following invariant
// holds for the mark, position, limit, and capacity values:
// 		0 <= mark <= position <= limit <= capacity
// A newly-created buffer always has a position of zero and a mark that is undefined. The initial limit
// may be zero, or it may be some other value that depends upon the type of the buffer and the manner
// in which it is constructed. Each element of a newly-allocated buffer is initialized to zero.
type WrappedByteBuffer struct {
	m_bytes *bytes.Buffer

	// The byte order is used when reading or writing multibyte values, and when creating buffers that are
	// views of this byte buffer. The order of a newly-created byte buffer is always BigEndian.
	m_bo binary.ByteOrder
}

func NewWrappedByteBuffer() *WrappedByteBuffer {
	buffer := &WrappedByteBuffer{
		m_bytes: bytes.NewBuffer([]byte{}),
		m_bo:    binary.BigEndian,
	}
	return buffer

}

func (this *WrappedByteBuffer) GetM_Bytes()  *bytes.Buffer{
	return this.m_bytes
}
func (this *WrappedByteBuffer) Reset()  {
	this.m_bytes.Reset()
}

func (this *WrappedByteBuffer) ReadableSize() int {
	return this.m_bytes.Len()
}

// Order retrieves this buffer's byte order.
// The byte order is used when reading or writing multibyte values, and when creating buffers that are
// views of this byte buffer. The order of a newly-created byte buffer is always binary.BigEndian.
func (this *WrappedByteBuffer) Order() binary.ByteOrder {
	return this.m_bo
}

// SetOrder modifies this buffer's byte order.
func (this *WrappedByteBuffer) SetOrder(bo binary.ByteOrder) {
	this.m_bo = bo
}

// GetBytes is a relative bulk get method.
//
// This method transfers bytes from this buffer into the given destination array. If there are fewer
// bytes remaining in the buffer than are required to satisfy the request, that is, if length > remaining(),
// then no bytes are transferred and a BufferUnderflowException is thrown.
//
// Otherwise, this method copies length bytes from this buffer into the given array, starting at the current
// position of this buffer and at the given offset in the array. The position of this buffer is then incremented
// by length.
//
// dst - The array into which bytes are to be written
// offset - The offset within the array of the first byte to be written; must be non-negative and no larger
// 		than dst.length
// length - The maximum number of bytes to be written to the given array; must be non-negative and no larger
// 		than dst.length - offset
func (this *WrappedByteBuffer) GetBytes(dst []byte, offset, length int) error {
	if offset < 0 || offset > cap(dst) {
		return errors.New("bytebuffer/GetBytes: Offset must be non-negative and no larger than length of dst")
	}

	if length < 0 || length > cap(dst)-offset {
		return errors.New("bytebuffer/GetBytes: Length must be non-negative and no larger than length of dst - offset")
	}

	if length > this.m_bytes.Len() {
		fmt.Println(length)
		fmt.Println(this.m_bytes.Len())
		return errors.New("bytebuffer/GetBytes: Insufficient bytes to get. Length is greater than remaining bytes.")
	}

	copy(dst[offset:], this.m_bytes.Next(length))

	return nil
}

// PutBytes is a relative bulk put method
//
// This method transfers bytes into this buffer from the given source array. If there are more bytes to be
// copied from the array than remain in this buffer, that is, if length > remaining(), then no bytes are
// transferred and a BufferOverflowException is thrown.
//
// Otherwise, this method copies length bytes from the given array into this buffer, starting at the given
// offset in the array and at the current position of this buffer. The position of this buffer is then
// incremented by length.
func (this *WrappedByteBuffer) PutBytes(src []byte, offset, length int) error {

	if offset < 0 || offset > cap(src) {
		return errors.New("bytebuffer/PutBytes: Offset must be non-negative and no larger than length of src")
	}

	if length < 0 || length > cap(src)-offset {
		return errors.New("bytebuffer/PutBytes: Length must be non-negative and no larger than length of src - offset")
	}
	this.m_bytes.Write(src[offset : offset+length])

	return nil
}
func (this *WrappedByteBuffer) Put(value byte) error {
	this.m_bytes.WriteByte(value)
	return nil
}

// GetUint16 is a relative get method for reading a short value.
// Reads the next two bytes at this buffer's current position, composing them into a short value
// according to the current byte order, and then increments the position by two.
func (this *WrappedByteBuffer) GetUint16() (uint16, error) {
	if this.m_bytes.Len() < 2 {
		return 0, errors.New("bytebuffer/Put: Insufficient remaining buffer for Uint16")
	}

	result := this.m_bo.Uint16(this.m_bytes.Next(2))
	return result, nil
}


// GetUint16 is a relative get method for reading a short value.
// Reads the next two bytes at this buffer's current position, composing them into a short value
// according to the current byte order, and then increments the position by two.
func (this *WrappedByteBuffer) PutUint16(value uint16) error {
	var buf []byte = make([]byte,2)
	this.m_bo.PutUint16(buf, value)
	this.m_bytes.Write(buf)
	return nil
}


// GetUint32 is a relative get method for reading a uint32 value.
// Reads the next four bytes at this buffer's current position, composing them into a uint32 value
// according to the current byte order, and then increments the position by two.
func (this *WrappedByteBuffer) GetUint32() (uint32, error) {
	if this.m_bytes.Len() < 4 {
		return 0, errors.New("bytebuffer/GetUint32: Insufficient remaining buffer for Uint32")
	}

	result := this.m_bo.Uint32(this.m_bytes.Next(4))
	return result, nil
}


// GetUint32 is a relative get method for reading a uint32 value.
// Reads the next four bytes at this buffer's current position, composing them into a uint32 value
// according to the current byte order, and then increments the position by two.
func (this *WrappedByteBuffer) PutUint32(value uint32) error {
	var buf []byte = make([]byte,4)
	this.m_bo.PutUint32(buf, value)
	this.m_bytes.Write(buf)
	return nil
}

// GetUint64 is a relative get method for reading a uint64 value.
// Reads the next eight bytes at this buffer's current position, composing them into a uint64 value
// according to the current byte order, and then increments the position by two.
func (this *WrappedByteBuffer) GetUint64() (uint64, error) {
	if this.m_bytes.Len() < 8 {
		return 0, errors.New("bytebuffer/Put: Insufficient remaining buffer for Uint64")
	}

	result := this.m_bo.Uint64(this.m_bytes.Next(8))
	return result, nil
}

// GetUint64 is a relative get method for reading a uint64 value.
// Reads the next eight bytes at this buffer's current position, composing them into a uint64 value
// according to the current byte order, and then increments the position by two.
func (this *WrappedByteBuffer) PutUint64(value uint64) error {
	var buf []byte = make([]byte,8)
	this.m_bo.PutUint64(buf, value)
	this.m_bytes.Write(buf)
	return nil
}


func (this *WrappedByteBuffer) String() string {
	return fmt.Sprintf("bytebuffer/String: length = %d, order = %s\n", this.m_bytes.Len(), this.m_bo.String())
}
