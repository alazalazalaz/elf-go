package traceid

import (
	"github.com/google/uuid"
)

var TraceId string

func GenTraceId() {
	TraceId = uuid.New().String()
}
