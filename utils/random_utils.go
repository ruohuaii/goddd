package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/sony/sonyflake"
)

var sf = sonyflake.NewSonyflake(sonyflake.Settings{
	StartTime: time.Date(2021, 11, 04, 0, 0, 0, 0, time.UTC),
	MachineID: func() (uint16, error) {
		return 1, nil
	},
})

func RandomString(length int) (string, error) {
	if length <= 0 {
		return "", errors.New("length must be greater than 0")
	}
	// 计算所需的字节数，确保生成的字符串长度为目标长度
	numBytes := (length + 1) / 2 // 每个字节编码为2个字符，如果长度是奇数，需要加一个字节
	b := make([]byte, numBytes)
	_, err := rand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate random bytes: %v", err)
	}
	return hex.EncodeToString(b)[:length], nil
}

func RandomNumber() (uint64, error) {
	if sf == nil {
		return 0, fmt.Errorf("failed to generate random number")
	}
	id, err := sf.NextID()
	if err != nil {
		return 0, fmt.Errorf("failed to generate random number: %v", err)
	}
	return id, nil
}
