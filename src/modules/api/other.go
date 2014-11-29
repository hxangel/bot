package api

type Other struct {
	ApiBase
}

func (c *Other) Get() {
	id := c.GetInput("url")
	callback := c.GetInput("callback")
	if id == "" {
		c.Json(-1, "with empty url", "")
		return
	}
	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}

	spiderServ.Add("Other", id, callback)
	c.Json(0, "success", "success")
}
