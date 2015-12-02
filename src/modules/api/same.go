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
	spiderServ.Add("same", map[string]string{"id": id, "channel":channel,"title":title, "callback": callback})
	c.Json(0, "success", "success")
}

func (c *Same) Shop() {
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

	spiderServ.Add("JdShop", map[string]string{"id": id, "callback": callback})
	c.Json(0, "success", "success")
}
