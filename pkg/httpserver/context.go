package httpserver

import "github.com/gin-gonic/gin"

type Context struct {
	*gin.Context
	isNext bool
}

func buildContext(c *gin.Context) *Context {
	return &Context{Context: c}
}

func (c *Context) Next() (RestfulResult, error) {
	c.Context.Next()
	c.isNext = true
	return nil, nil
}

func (c *Context) isNexted() bool {
	return c.isNext
}
