package gin

import (
	"github.com/gentwolf-shen/gin-boost/binding"
)

func (c *Context) BindRequest(obj interface{}) error {
	m := make(map[string][]string)
	for _, v := range c.Params {
		m[v.Key] = []string{v.Value}
	}

	if err := binding.Simple.BindExt(c.Request, obj, m); err != nil {
		return err
	}

	method := c.Request.Method
	if method == "POST" || method == "PUT" || method == "PATCH" {
		return c.ShouldBind(obj)
	}

	return nil
}
