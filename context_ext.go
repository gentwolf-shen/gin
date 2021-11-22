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

func (c *Context) AllowCrossDomainAll() HandlerFunc {
	return func(c *Context) {
		c.allowCrossDomain("*")
	}
}

func (c *Context) AllowCrossDomain(domains []string) HandlerFunc {
	return func(c *Context) {
		host := c.Request.Header.Get("Origin")
		bl := false
		for _, domain := range domains {
			if domain == host {
				bl = true
				break
			}
		}

		if bl {
			c.allowCrossDomain(c.Request.Header.Get("Origin"))
		}
	}
}

func (c *Context) allowCrossDomain(host string) {
	c.Header("Access-Control-Allow-Origin", host)
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, X-Requested-With, Content-Type, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATH,DELETE,OPTIONS,HEAD")
		c.Header("Access-Control-Max-Age", "3600")
		c.AbortWithStatus(200)
	}
}
