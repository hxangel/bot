package api

import (
	"net/url"
)

type Other struct {
	ApiBase
}

func (c *Other) Get() {
	urlStr := c.GetInput("url")
	callback := c.GetInput("callback")
	if urlStr == "" {
		c.Json(-1, "with empty url", "")
		return
	}
	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}
	link, _ := url.QueryUnescape(urlStr)
	spiderServ.Add("Other", map[string]string{"link": link, "callback": callback})
	c.Json(0, "success", "success")
}
