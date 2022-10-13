package enum

import "elf-go/components/apphelper"

const (
	RespInvalidParams apphelper.ErrorCode = 1001
	RespDbError       apphelper.ErrorCode = 1002
	RespReadFileError apphelper.ErrorCode = 1003
	RespGenJWTError   apphelper.ErrorCode = 1004 //生成JWT错误

	RespSystemError apphelper.ErrorCode = 500
)
