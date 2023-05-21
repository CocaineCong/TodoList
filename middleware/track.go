package middleware

import (
	"to-do-list/pkg/track"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
)

var (
	SpanCTX = "span-ctx"
)

func Jaeger() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("uber-trace-id")
		var span opentracing.Span
		if traceId != "" {
			var err error
			span, err = track.GetParentSpan(c.FullPath(), traceId, c.Request.Header)
			if err != nil {
				return
			}
		} else {
			span = track.StartSpan(opentracing.GlobalTracer(), c.FullPath())
		}
		defer span.Finish()

		c.Set(SpanCTX, opentracing.ContextWithSpan(c, span))
		c.Next()
	}
}
