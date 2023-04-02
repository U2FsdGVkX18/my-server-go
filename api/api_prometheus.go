package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"time"
)

func Prometheus(ginServer *gin.Engine) {
	//接入Prometheus
	reqCounter := prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "Number of HTTP requests",
	}, []string{"method", "endpoint", "status"})

	//请求持续时间
	reqDuration := prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "http_request_duration_seconds",
		Help:    "Duration of HTTP requests",
		Buckets: prometheus.DefBuckets,
	}, []string{"method", "endpoint", "status"})

	prometheus.MustRegister(reqCounter, reqDuration)

	//使用中间件
	ginServer.Use(func(context *gin.Context) {
		method := context.Request.Method
		endpoint := context.FullPath()

		start := time.Now()

		context.Next()

		duration := time.Since(start)
		status := context.Writer.Status()

		reqCounter.WithLabelValues(method, endpoint, fmt.Sprintf("%d", status)).Inc()
		reqDuration.WithLabelValues(method, endpoint, fmt.Sprintf("%d", status)).Observe(duration.Seconds())
	})

	//Add Prometheus metrics endpoint
	ginServer.GET("/metrics", gin.WrapH(promhttp.Handler()))
}
