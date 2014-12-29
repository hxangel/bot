package api

import (
	"fmt"
	spider "github.com/rainkid/spider"
	"time"
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
	callback := c.GetInput("callback")
	title := c.GetInput("title")
	if id == "" {
		c.Json(-1, "with empty id", "")
		return
	}
	if title == "" {
		c.Json(-1, "with empty title", "")
		return
	}
	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}
	pid := c.getUnipid(id, title)
	if pid != nil {
		spiderServ.Add("SameStyle", map[string]string{"callback": callback, "id": id, "pid": string(pid)})
		c.Json(0, "success", string(pid))
		return
	}
	c.Json(-1, "fail", nil)
	return
}

func (c *Taobao) getUnipid(id, title string) []byte {
	var (
		pid           []byte
		precessed     int = 0   //已经完成进程数
		processNum    int = 3   //开启进程数
		timeOut       int = 600 //接口超时设置
		dataPrecessNo int = 0   //获取到进程的进程编号
	)
	Loger.D("Get unipid for", id, title)
	for i := 1; i <= processNum; i++ {
		go func(i int) {
			surl := fmt.Sprintf("http://s.taobao.com/search?q=%s", title)
			sloader := spider.NewLoader(surl, "Get").WithPcAgent()
			scontent, _ := sloader.Send(nil)

			if pid != nil {
				Loger.W(fmt.Sprintf("The pid-%d process cancel to parse data", i))
				return
			}
			shp := spider.NewHtmlParse().LoadData(scontent).Replace().Convert()
			ret := shp.Partten(`(?U)"samestyle":{"url":".*uniqpid=-(\d+)&nid=` + id + `"}`).FindStringSubmatch()
			if ret != nil && len(ret) > 0 {
				pid = ret[1]
				dataPrecessNo = i
			}
			Loger.I("The process finished is pid-", i)
			precessed += 1
			return
		}(i)
	}
	var wait int = 0
	for {
		if pid != nil {
			Loger.I(fmt.Sprintf("The pid-%d process finished and get unipid :%s.", dataPrecessNo, pid))
			break
		}
		//超过5秒
		if wait == timeOut {
			Loger.E("Get unipid timeout with item id", id)
			break
		}
		//如果进程都已经完成并且没有拉取到数据
		if precessed == processNum {
			Loger.E(fmt.Sprintf("%d process finished and not found unipid.", processNum))
			break
		}
		time.Sleep(time.Second / 60)
		wait++
	}
	return pid
}
