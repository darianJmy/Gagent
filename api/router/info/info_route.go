package info

import (
	"github.com/emicklei/go-restful/v3"

	"gagent/api/middleware"
)

func Register(container *restful.Container) {
	ws := new(restful.WebService)

	ws.Filter(middleware.NCSACommonLogFormatLogger())
	ws.Path("/info").
		Consumes(restful.MIME_XML, restful.MIME_JSON).
		Produces(restful.MIME_XML, restful.MIME_JSON)

	ws.Route(ws.GET("/cpu").To(GetCPUInfo))
	ws.Route(ws.GET("/network").To(GetNetWorkInfo))
	ws.Route(ws.GET("/system/uuid").To(GetSystemUUID))

	container.Add(ws)
}
