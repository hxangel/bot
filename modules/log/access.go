package log

import (
	"github.com/hxangel/dogo"

)

type Access struct {
	dogo.Controller
}

func (c *Access) Index() {
	c.Assign("var","Demo")
	c.Assign("content","Demo")
	c.Render();
}


