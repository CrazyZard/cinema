package request

import (
	"cinema/model/common/request"
	"time"
)

type CinemaOrderSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
	OrderStatus    int        `json:"orderStatus" form:"orderStatus"`
	request.PageInfo
}
