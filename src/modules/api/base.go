package api

import (
	// "fmt"
	"github.com/rainkid/dogo"
	spider "github.com/rainkid/spider"
)

var (
	spiderServ *spider.Spider
	Loger      = dogo.NewLoger()
)

type ApiBase struct {
	dogo.Controller
}

func (c *ApiBase) Init() {
	c.DisableView = true
	spiderServ = spider.SpiderServer
}
