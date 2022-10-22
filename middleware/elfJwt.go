package middleware

import (
	"elf-go/components/appconsts"
	"elf-go/components/apphelper"
	"elf-go/components/appjwts"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

func ParseJwt(ctx *gin.Context) {
	auth := ctx.Request.Header.Get(appconsts.HeaderAuthorization)
	if auth == "" {
		apphelper.EchoFailed(ctx, apphelper.StatusUnauthorized, "empty Authorization header")
		ctx.Abort()
		return
	}

	if !strings.HasPrefix(auth, "Bearer ") {
		apphelper.EchoFailed(ctx, apphelper.StatusUnauthorized, "Authorization header has no Bearer prefix")
		ctx.Abort()
		return
	}

	token := auth[7:]
	if err := appjwts.ParseJwtToken(token); err != nil {
		apphelper.EchoFailed(ctx, apphelper.StatusUnauthorized, fmt.Sprintf("parse JWT token error:%v", err))
		ctx.Abort()
		return
	}

	ctx.Next()
}
