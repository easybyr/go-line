package util

import (
	"crypto/md5"
	"fmt"
)

func MD5(str string) string {
	data := []byte(str)
	checkSum := md5.Sum(data)
	return fmt.Sprintf("%x", checkSum)
}
