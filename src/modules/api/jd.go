package api

type Jd struct {
	ApiBase
}

func (c *Jd) Item() {
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

	spiderServ.Add("JdItem", map[string]string{"id": id, "callback": callback})
	c.Json(0, "success", "success")
}
func (c *Jd) Itemhk() {
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

	spiderServ.Add("JdItemHk", map[string]string{"id": id, "callback": callback})
	c.Json(0, "success", "success")
}

func (c *Jd) Shop() {
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
func (c *Jd) Shophk() {
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

	spiderServ.Add("JdShopHk", map[string]string{"id": id, "callback": callback})
	c.Json(0, "success", "success")
}
