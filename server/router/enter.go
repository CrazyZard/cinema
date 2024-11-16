package router

import (
	"cinema/router/cinema"
	"cinema/router/example"
	"cinema/router/system"
)

var RouterGroupApp = new(RouterGroup)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Cinema  cinema.RouterGroup
}
