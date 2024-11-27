package trace

import (
	"context"
	"github.com/gin-gonic/gin"
)

func TraceId() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "__traceId__", c.GetHeader("traceId")))
	}
}
