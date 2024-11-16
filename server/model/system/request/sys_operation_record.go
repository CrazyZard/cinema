package request

import (
	"cinema/model/common/request"
	"cinema/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
