package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

// convert Int64 to []256-bits
func IntToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)
	}

	return buff.Bytes()
}
