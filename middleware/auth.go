package middleware

import (
	"crypto/md5"
	"fmt"
	"github.com/gin-gonic/gin"
	"slices"
)

// SecurityVerification 身份验证
func SecurityVerification(c *gin.Context, k []string) bool {
	XAPIKey := c.GetHeader("X-API-Key")
	return AuthKey(k, XAPIKey)

}

// AuthKey 加密并验证是否存在
func AuthKey(keylocal []string, XAPIKey string) bool {
	XAPIKey_md5 := fmt.Sprintf("%x", md5.Sum([]byte(XAPIKey)))
	return slices.Contains(keylocal, XAPIKey_md5)
}
