package transport

import (
	"github.com/gin-gonic/gin"
)

func RegisterMetricsEndpoint(router *gin.Engine) {
	metricsGroup := router.Group("/metrics")
	{
		metricsGroup.GET("", prometheusHandler())
	}
}
