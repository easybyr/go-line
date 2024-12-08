package util

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)


func GenUUID() string {
	uuid, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	return strings.ReplaceAll(uuid.String(), "-", "")
}

func GenPUUID(prefix string) string {
	return fmt.Sprintf("%s%s", prefix, GenUUID())
}


