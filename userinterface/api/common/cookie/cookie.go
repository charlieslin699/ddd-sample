package cookie

import "ddd-sample/pkg/httpserver"

var (
	// JWT token
	AuthToken Cookie = newCookie("AuthToken")
	// 語系
	Locale Cookie = newCookie("Locale",
		WithDefaultValue("zh-tw"),
	)
)

type Cookie struct {
	name         string // Cookie name
	defaultValue string // 預設值
	expires      int    // 到期時間(單位:秒), 0為無限期
	path         string
	domain       string
	secure       bool // 是否限https
	httpOnly     bool // 是否禁止Javascript存取
}

func newCookie(name string, fns ...CookieOptionFunc) Cookie {
	c := Cookie{
		name:     name,
		expires:  86400, // 預設一天到期
		httpOnly: true,
	}

	for _, fn := range fns {
		fn(&c)
	}

	return c
}

func (c Cookie) Get(ctx *httpserver.Context) string {
	value, err := ctx.Cookie(c.name)
	if err != nil {
		return c.defaultValue
	}

	return value
}

func (c Cookie) Set(ctx *httpserver.Context, value string, fns ...CookieOptionFunc) {
	copyCookie := c

	for _, fn := range fns {
		fn(&copyCookie)
	}

	ctx.SetCookie(
		copyCookie.name,
		value,
		copyCookie.expires,
		copyCookie.path,
		copyCookie.domain,
		copyCookie.secure,
		copyCookie.httpOnly)
}

type CookieOptionFunc func(c *Cookie)

func WithDefaultValue(defaultValue string) CookieOptionFunc {
	return func(c *Cookie) {
		c.defaultValue = defaultValue
	}
}

func WithExpires(second int) CookieOptionFunc {
	return func(c *Cookie) {
		c.expires = second
	}
}

func WithPath(path string) CookieOptionFunc {
	return func(c *Cookie) {
		c.path = path
	}
}

func WithDomain(domain string) CookieOptionFunc {
	return func(c *Cookie) {
		c.domain = domain
	}
}

func WithSecure(secure bool) CookieOptionFunc {
	return func(c *Cookie) {
		c.secure = secure
	}
}

func WithHTTPOnly(httpOnly bool) CookieOptionFunc {
	return func(c *Cookie) {
		c.httpOnly = httpOnly
	}
}
