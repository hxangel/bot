package api

import (
	"fmt"
	spider "github.com/rainkid/spider"
	"net/http"
	"io/ioutil"
)

type Taobao struct {
	ApiBase
}

func (c *Taobao) Item() {
	id := c.GetInput("id")
	callback := c.GetInput("callback")
	if id == "" {
		c.Json(-1, "with empty id", "")
		return
	}
	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}

	spiderServ.Add("TaobaoItem", map[string]string{"id": id, "callback": callback})
	c.Json(0, "success", "success")
}

func (c *Taobao) Shop() {
	id := c.GetInput("id")
	callback := c.GetInput("callback")
	if id == "" {
		c.Json(-1, "with empty id", "")
		return
	}
	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}

	spiderServ.Add("TaobaoShop", map[string]string{"id": id, "callback": callback})
	c.Json(0, "success", "success")
}

func (c *Taobao) Samestyle() {
	id := c.GetInput("id")
	pid := c.GetInput("pid")
	callback := c.GetInput("callback")
	title := c.GetInput("title")
	if id == "" {
		c.Json(-1, "with empty id", "")
		return
	}

	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}

	go c.getUnipid(pid, id, title, callback)

	c.Json(-1, "success", nil)
	return
}

func (c *Taobao) getUnipid(pid, id, title, callback string)  {
	if pid == "" ||pid =="0"{
		surl := fmt.Sprintf("http://s.taobao.com/search?q=%s", title)
		resp, err := http.Get(surl)
		if err != nil {
			// handle error
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			Loger.I("The process get error",)
			return
		}
		mcontent :=make([]byte,len(body))
		copy(mcontent, body)

		shp := spider.NewHtmlParser().LoadData(mcontent)

		ret := shp.Partten(`(?U)"nid":"`+id+`","category":"\d+","pid":"-(\d+)"`).FindStringSubmatch()
		if ret != nil && len(ret) > 0 {
			pid = string(ret[1])
		}
	}
	if pid == "" ||pid =="0"{
		return
	}
	spiderServ.Add("SameStyle", map[string]string{"callback": callback, "id": id, "pid": pid})
	return
}
