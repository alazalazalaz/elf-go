package middleware

import (
	"elf-go/components/apphelper"
	"elf-go/components/applogs"
	"errors"
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

const (
	SLIDE_WINDOWS_LIMIT      = 20               //限流值是20
	SLIDE_WINDOWS_LIMIT_UNIT = 60 * time.Second //限流单位是每分钟
)

var slideWindows struct {
	Qps  int64
	Pool sync.Map
}

func Limit(ctx *gin.Context) {
	//获取请求时间
	reqAt := time.Now().Unix()

	//判断是否被限流
	limited, err := isLimited(reqAt)
	if err != nil {
		applogs.Errorf("limit error:%v, at:%v", err, time.Now().Format(time.RFC3339))
		ctx.Next()
		return
	}

	//如果是 ，则返回失败
	if limited {
		applogs.Warnf("request limited")
		apphelper.EchoFailed(ctx, apphelper.StatusRequestLimited, "request limited")
		ctx.Abort()
		return
	}

	ctx.Next()
}

func isLimited(reqAt int64) (bool, error) {
	limitUnitFrom := reqAt - int64(SLIDE_WINDOWS_LIMIT_UNIT.Seconds())
	var qps int64

	// 把过去的时间单位删除
	slideWindows.Pool.Range(func(key, value interface{}) bool {
		var keyInt64, valueInt64 int64
		var ok bool
		if keyInt64, ok = key.(int64); !ok {
			return false
		}
		if valueInt64, ok = value.(int64); !ok {
			return false
		}

		applogs.Infof("limit pool, key=%v, value=%v", time.Unix(keyInt64, 0).Format(time.RFC3339), valueInt64)
		if keyInt64 < limitUnitFrom {
			slideWindows.Pool.Delete(key)
		} else {
			qps += valueInt64
		}

		return true
	})

	// 计算限流值
	if qps > SLIDE_WINDOWS_LIMIT {
		return true, nil
	}

	// 把当前的时间单位新增上来
	var currentUnitValue int64
	if value, ok := slideWindows.Pool.Load(reqAt); ok {
		var isOk bool
		if currentUnitValue, isOk = value.(int64); !isOk {
			return false, errors.New("slide windows limit failed, value.(int64) failed")
		}
	}

	slideWindows.Pool.Store(reqAt, currentUnitValue+1)
	return false, nil
}
