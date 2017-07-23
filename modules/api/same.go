package api

type Same struct {
	ApiBase
}

func (c *Same) Item() {
	id       := c.GetInput("id")
	channel  := c.GetInput("channel")
	title    := c.GetInput("title")
	callback := c.GetInput("callback")
	if id == "" {
		c.Json(-1, "with empty id", "")
		return
	}
	if callback == "" {
		c.Json(-1, "with empty callback", "")
		return
	}
	if channel == "" {
		c.Json(-1, "with empty channel", "")
		return
	}
	if title == "" {
		c.Json(-1, "with empty title", "")
		return
	}
	go func() {
		spiderServ.Add("Same", map[string]string{"id": id, "channel":channel,"title":title, "callback": callback})
	}()
	c.Json(0, "success", "success")
}
