package gin

import (
	"github.com/gentwolf-shen/gin/binding"
)

func (c *Context) BindRequest(obj interface{}) error {
	if err := c.MustBindWith(obj, binding.Simple); err != nil {
		return err
	}

	method := c.Request.Method
	if method == "POST" || method == "PUT" || method == "PATCH" {
		return c.ShouldBind(obj)
	}

	return nil
}
