package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
	"weatherAPI/src/lib/log"
)

func AccessLog(c *gin.Context) {
	n := time.Now()
	c.Next()
	da, _ := c.GetRawData()
	log.WithContext(c).With(
		"ua", c.Request.UserAgent(),
		"referer", c.Request.Referer(),
		"data", string(da)).
		Infof("%s %d %s %s %s %v",
			c.ClientIP(),
			c.Writer.Status(),
			c.Request.Method,
			c.Request.Proto,
			c.Request.RequestURI,
			time.Now().Sub(n).String(),
		)
}