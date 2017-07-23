package api

type Mmb struct {
	ApiBase
}

func (c *Mmb) Item() {
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
	spiderServ.Add("MmbItem", map[string]string{"id": id, "callback": callback})
	c.Json(0, "success", "success")
}
