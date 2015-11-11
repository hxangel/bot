package configs

import (
	"github.com/rainkid/dogo"
	api "modules/api"
	debug "modules/debug"
)

func AddSampleRoute(router *dogo.Router) {
	router.AddSampleRoute("api", &api.Jd{})
	router.AddSampleRoute("api", &api.Mmb{})
	router.AddSampleRoute("api", &api.Other{})
	router.AddSampleRoute("api", &api.Taobao{})
	router.AddSampleRoute("api", &api.Tmall{})
	router.AddSampleRoute("debug", &debug.Pprof{})
}