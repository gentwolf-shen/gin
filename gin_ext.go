package gin

func AllowCrossDomainAll() HandlerFunc {
	return func(c *Context) {
		allowCrossDomain(c, "*")
	}
}

func AllowCrossDomain(domains []string) HandlerFunc {
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
			allowCrossDomain(c, c.Request.Header.Get("Origin"))
		}
	}
}

func allowCrossDomain(c *Context, host string) {
	c.Header("Access-Control-Allow-Origin", host)
	c.Header("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, X-Requested-With, Content-Type, Authorization")
	c.Header("Access-Control-Allow-Credentials", "true")
	if c.Request.Method == "OPTIONS" {
		c.Header("Access-Control-Allow-Methods", "POST,GET,PUT,DELETE,OPTIONS")
		c.Header("Access-Control-Max-Age", "3600")
		c.AbortWithStatus(200)
	}
}
