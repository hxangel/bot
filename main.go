package main

import (
	"flag"
	"github.com/rainkid/dogo"
	spider "github.com/rainkid/spider"
	"os"
	"path"
	"runtime"
)

var (
	host = flag.String("h", "0.0.0.0", "hostname for app runtime")
	port = flag.String("p", "8090", "port for app runtime")
	Loger  = dogo.NewLoger()
)

func main() {
	flag.Parse()

	defer func() {
		if err := recover(); err != nil {
			Loger.E("run time panic: ", err)
		}
	}()

	//start spider daemon and  proxy daemon
	spider.Start()
	spider.StartProxy()

	router := getRouter()

	runtime.GOMAXPROCS(runtime.NumCPU() - 1)

	// bootstrap and return a app
	app := dogo.NewApp(*host, *port)
	//Bootstrap and run
	app.Bootstrap(router).SetDefaultModule("api").Run()
}

func getRouter() *dogo.Router {
	// new dogo router
	var router = dogo.NewRouter()
	basepath, _ := os.Getwd()

	//add static router
	router.AddStaticRoute("/statics", path.Join(basepath, "src/statics/"))

	//add sample route
	AddSampleRoute(router)
	return router
}
