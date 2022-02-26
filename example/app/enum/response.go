package enum

import "elf-go/utils/helper"

const (
	RespInvalidParams helper.ErrorCode = 1001
	RespDbError       helper.ErrorCode = 1002
	RespReadFileError helper.ErrorCode = 1003

	RespSystemError helper.ErrorCode = 500
)
