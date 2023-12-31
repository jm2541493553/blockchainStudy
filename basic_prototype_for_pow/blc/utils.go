package blc

import (
	"bytes"
	"encoding/binary"
	"log"
)

// IntToHex 将int64 转成 []byte
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}
	return buff.Bytes()
}
