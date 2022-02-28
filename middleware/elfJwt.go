package middleware

import (
	"elf-go/components/jwts"
	"elf-go/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ParseJwt(ctx *gin.Context) {
	auth := ctx.Request.Header.Get("Authorization")
	if auth == "" {
		helper.EchoFailed(ctx, helper.StatusUnauthorized, "empty Authorization header")
		ctx.Abort()
		return
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		helper.EchoFailed(ctx, helper.StatusUnauthorized, "Authorization header has no Bearer prefix")
		ctx.Abort()
		return
	}

	token := auth[7:]
	if err := jwts.ParseJwtToken(token); err != nil {
		helper.EchoFailed(ctx, helper.StatusUnauthorized, fmt.Sprintf("parse JWT token error:%v", err))
		ctx.Abort()
		return
	}

	ctx.Next()
}
