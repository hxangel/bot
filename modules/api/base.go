package api

import (
	// "fmt"
	"github.com/hxangel/dogo"
	spider "github.com/hxangel/spider"
	"fmt"
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
	req := c.Context.Request
	url := fmt.Sprintf("Request [http://%s%s]", req.Host, req.URL)
	Loger.I(url)
}
