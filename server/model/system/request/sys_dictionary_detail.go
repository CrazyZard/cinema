package request

import (
	"cinema/model/common/request"
	"cinema/model/system"
)

type SysDictionaryDetailSearch struct {
	system.SysDictionaryDetail
	request.PageInfo
}
