package service

import (
	"cinema/service/cinema"
	"cinema/service/example"
	"cinema/service/system"
)

var ServiceGroupApp = new(ServiceGroup)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	CinemaServiceGroup  cinema.ServiceGroup
}
