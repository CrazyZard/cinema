package request

import (
	"cinema/model/common/request"
	"time"
)

type CinemaFilmSearch struct {
	StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
	EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`

	Hall *int   `json:"hall" form:"hall" `
	Name string `json:"name" form:"name" `
	request.PageInfo
}
