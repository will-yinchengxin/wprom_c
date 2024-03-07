package test

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	wprom "github.com/will-yinchengxin/wprom_c/core"
	"strconv"
	"testing"
)

func TestWpromC(t *testing.T) {
	gin.SetMode("release")
	r := gin.New()
	reqMetric := wprom.SetGaugeVecMetric(
		"will_test",
		"requests_total",
		"How many HTTP requests processed, partitioned by status code and HTTP method.",
		[]string{"code", "method", "handler", "host", "path"},
	)
	p := wprom.NewPrometheus(reqMetric).
		SetSubsystemAndRegisterMetrics("will")
	p.Use(r, HandleFun(p))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello world!")
	})

	r.Run(":29090")
}

func HandleFun(p *wprom.Prometheus) gin.HandlerFunc {
	return func(c *gin.Context) {
		for key, _ := range p.MetricsList {
			c.Next()
			if key == "will_test" {
				status := strconv.Itoa(c.Writer.Status())
				p.MetricsList[key].MetricCollector.(*prometheus.GaugeVec).
					WithLabelValues(status, c.Request.Method, c.HandlerName(), c.Request.Host, c.Request.URL.Path).
					Inc()
			}
		}
	}
}
