// Package trader package trader
package trader

import (
	"tmt/internal/usecase/modules/cache"
	"tmt/internal/usecase/modules/event"

	"tmt/internal/usecase/modules/logger"
)

var (
	cc  = cache.Get()
	bus = event.Get()
	log = logger.Get()
)
