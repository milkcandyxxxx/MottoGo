package middleware

import (
	"crypto/md5"
	"fmt"
	"slices"

	"github.com/gin-gonic/gin"
)

// Security_verification 身份验证
func Security_verification(c *gin.Context, k []string) bool {
	authHeader := c.GetHeader("Authorization")
	return Auth_key(k, authHeader)

}

// Auth_key 加密并验证是否存在
func Auth_key(keylocal []string, Authorization string) bool {
	Authorization_md5 := fmt.Sprintf("%x", md5.Sum([]byte(Authorization)))
	return slices.Contains(keylocal, Authorization_md5)
}
