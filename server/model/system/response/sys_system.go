package response

import "cinema/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
