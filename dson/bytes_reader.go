package dson

import (
	"bytes"
	"encoding/binary"
)

type BytesReader struct {
	bytes.Reader
}

func NewBytesReader(bs []byte) BytesReader {
	return BytesReader{
		Reader: *bytes.NewReader(bs),
	}
}

func (b *BytesReader) ReadInt() (int, error) {
	bs := make([]byte, 4)
	_, err := b.Read(bs)
	if err != nil {
		return 0, err
	}
	result := binary.LittleEndian.Uint32(bs)
	return int(int32(result)), nil
}

func (b *BytesReader) ReadLong() (int64, error) {
	bs := make([]byte, 8)
	_, err := b.Read(bs)
	if err != nil {
		return 0, err
	}
	result := binary.LittleEndian.Uint64(bs)
	return int64(result), nil
}

func (b *BytesReader) ReadBytes(n int) ([]byte, error) {
	bs := make([]byte, n)
	_, err := b.Read(bs)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func (b *BytesReader) ReadString(n int) (string, error) {
	bs, err := b.ReadBytes(n)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
