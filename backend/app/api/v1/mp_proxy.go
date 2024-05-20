package v1

import (
	"github.com/1Panel-dev/1Panel/backend/global"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"net/url"
	"strings"
)

func (b *BaseApi) ApiProxyHandler(c *gin.Context) {
	hostId := c.GetHeader("host-id")
	if strings.Contains(c.GetHeader("Upgrade"), "websocket") {
		if len(c.Query("host-id")) > 0 {
			hostId = c.Query("host-id")
		}
	}
	token, err := global.CACHE.Get("mp-token:" + hostId)
	if err != nil {
		return
	}
	host, err := global.CACHE.Get("mp-host:" + hostId)
	if err != nil {
		return
	}
	c.Request.Header.Set("Panelauthorization", string(token))
	proxyUrl, _ := url.Parse(string(host))
	proxy := httputil.NewSingleHostReverseProxy(proxyUrl)
	proxy.ServeHTTP(c.Writer, c.Request)
}
