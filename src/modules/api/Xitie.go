package api
import "fmt"

type Xitie struct {
	ApiBase
}

func (c *Xitie) Item() {
	id       := c.GetInput("id")
	channel  := c.GetInput("channel")
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
	fmt.Println(channel)
	spiderServ.Add("XtItem", map[string]string{"id": id, "channel":channel, "callback": callback})
	c.Json(0, "success", "success")
}

func (c *Xitie) Shop() {
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
