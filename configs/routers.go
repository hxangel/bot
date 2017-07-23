package configs

import (
	"git-ss.gionee.com/Go/dogo"
	debug "github.com/hxangel/bot/modules/debug"
	log "github.com/hxangel/bot/modules/log"
	"github.com/hxangel/bot/modules/front"
	"github.com/hxangel/bot/modules/admin"
	"github.com/hxangel/bot/modules/api"
)

func AddSampleRoute(router *dogo.Router) {
	router.AddSampleRoute("api", &api.Jd{})
	router.AddSampleRoute("api", &api.Mmb{})
	router.AddSampleRoute("api", &api.Same{})
	router.AddSampleRoute("api", &api.Other{})
	router.AddSampleRoute("api", &api.Taobao{})
	router.AddSampleRoute("api", &api.Tmall{})
	router.AddSampleRoute("log", &log.Access{})
	router.AddSampleRoute("admin", &admin.Errors{})
	router.AddSampleRoute("admin", &admin.Group{})
	router.AddSampleRoute("admin", &admin.Index{})
	router.AddSampleRoute("admin", &admin.Login{})
	router.AddSampleRoute("admin", &admin.User{})
	router.AddSampleRoute("debug", &debug.Pprof{})
	router.AddSampleRoute("front", &front.Index{})
}