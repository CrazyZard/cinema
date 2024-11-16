package v1

import (
	"cinema/api/v1/cinema"
	"cinema/api/v1/example"
	"cinema/api/v1/system"
)

var ApiGroupApp = new(ApiGroup)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	CinemaApiGroup  cinema.ApiGroup
}
