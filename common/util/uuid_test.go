package util_test

import (
	"fmt"
	"testing"

	"github.com/easybyr/go-line/common/util"
)


func TestUUID(t *testing.T) {
	uuid := util.GenUUID()
	fmt.Println("uuid: " + uuid)

	uuid2 := util.GenPUUID("JO")
	fmt.Println("uuid prefix: " + uuid2)
}


func TestMd5(t *testing.T) {
	md5 := util.MD5("abc")
	fmt.Println("md5: " + md5)
}