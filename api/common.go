package api

import (
	"github.com/emicklei/go-restful/v3"

	"gagent/api/router/info"
)

func RegisterRoutes(wsContainer *restful.Container) {
	info.Register(wsContainer)
}
