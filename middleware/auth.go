package middleware

import (
	"crypto/md5"
	"fmt"
	"slices"
	"strings"
)

// AuthKey 加密并验证是否存在
func AuthKey(keylocal []string, authorization string) bool {
	const prefix = "Bearer "
	token := strings.TrimPrefix(authorization, prefix)
	authorization_md5 := fmt.Sprintf("%x", md5.Sum([]byte(token)))
	return slices.Contains(keylocal, authorization_md5)
}
