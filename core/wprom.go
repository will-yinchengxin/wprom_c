package core

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

var (
	defaultMetricPath    = "/metrics"
	defaultPromSubsystem = "wprom"
)

type Prometheus struct {
	router *gin.Engine

	listenAddress string

	PromSubsystem string

	MetricsList map[string]*Metric
	MetricsPath string

	URLLabelFromContext string
}

func NewPrometheus(customMetricsList ...*Metric) *Prometheus {
	var metricsList = map[string]*Metric{}

	if len(customMetricsList) < 0 {
		panic("No *Metric Send")
	}

	for _, metric := range customMetricsList {
		metricsList[metric.Id] = metric
	}

	p := &Prometheus{
		MetricsList: metricsList,
		MetricsPath: defaultMetricPath,
	}

	return p
}

func (p *Prometheus) RegisterMetrics() *Prometheus {
	p.registerMetrics()
	return p
}

func (p *Prometheus) SetSubsystemAndRegisterMetrics(j string) *Prometheus {
	p.PromSubsystem = j
	p.registerMetrics()
	return p
}

func (p *Prometheus) registerMetrics() {
	if p.PromSubsystem == "" {
		p.PromSubsystem = defaultPromSubsystem
	}
	for _, metricDef := range p.MetricsList {
		metric := NewMetric(p.PromSubsystem, metricDef)
		if err := prometheus.Register(metric); err != nil {
			log.WithError(err).Errorf("「%s」 could not be registered in Prometheus", metricDef.Name)
		}
		metricDef.MetricCollector = metric
	}
}

func (p *Prometheus) Use(e *gin.Engine, handelFun gin.HandlerFunc) {
	e.Use(handelFun)
	p.SetMetricsPath(e)
}

func (p *Prometheus) SetMetricsPath(e *gin.Engine) {
	if p.listenAddress != "" {
		p.router.GET(p.MetricsPath, prometheusHandler())
		p.runServer()
	} else {
		e.GET(p.MetricsPath, prometheusHandler())
	}
}

func prometheusHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		promhttp.Handler().ServeHTTP(c.Writer, c.Request)
	}
}

func (p *Prometheus) runServer() {
	if p.listenAddress != "" {
		go p.router.Run(p.listenAddress)
	}
}

func (p *Prometheus) SetListenAddress(address string) *Prometheus {
	p.listenAddress = address
	if p.listenAddress != "" {
		p.router = gin.Default()
	}
	return p
}
