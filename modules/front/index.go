package front

import (
	"fmt"
	// "math/rand"
	ws "github.com/hxangel/bot/libs/websock"
	"os"
	// "time"
	"github.com/hxangel/bot/libs/pserver"
)

var page int = 614
var total int = 0
var f *os.File

type Index struct {
	FrontBase
}

func (c *Index) Index() {

}


func (c *Index) Ppt() {

}

func (c *Index) Ppt1() {

}


func (c *Index) Next() {
	ws.Server.SendData([]byte("next"))
	c.Json(0,"success","")
}

func (c *Index) Server() {
	code := c.GetInput("code")
	pserver.Server.SendData([]byte(code))
}

func (c *Index) Test() {
	c.DisableView = true
	params := c.GetPosts([]string{"payStatus", "queryType", "startTime", "endTime", "cookie"})
	fmt.Println(params)
	c.Json(0, "request is submit, please wait.", "")
}


