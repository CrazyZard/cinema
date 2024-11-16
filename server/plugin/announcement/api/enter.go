package api

import "cinema/plugin/announcement/service"

var (
	Api         = new(api)
	serviceInfo = service.Service.Info
)

type api struct{ Info info }
