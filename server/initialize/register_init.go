package initialize

import (
	_ "cinema/source/example"
	_ "cinema/source/system"
)

func init() {
	// do nothing,only import source package so that inits can be registered
}
