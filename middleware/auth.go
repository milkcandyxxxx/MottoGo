package middleware

import (
	"crypto/md5"
	"fmt"
	"slices"

	"github.com/gin-gonic/gin"
)

// SecurityVerification 身份验证
func SecurityVerification(c *gin.Context, k []string) bool {
	authHeader := c.GetHeader("Authorization")
	return AuthKey(k, authHeader)

}

// AuthKey 加密并验证是否存在
func AuthKey(keylocal []string, Authorization string) bool {
	Authorization_md5 := fmt.Sprintf("%x", md5.Sum([]byte(Authorization)))
	return slices.Contains(keylocal, Authorization_md5)
}
