package traceid

import (
	"github.com/google/uuid"
)

var TraceId string
// 这样的trace id会有并发问题，因为是全局变量
func GenTraceId() {
	TraceId = uuid.New().String()
}
