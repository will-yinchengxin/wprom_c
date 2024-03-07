package core

import "github.com/prometheus/client_golang/prometheus"

const (
	CounterVec   = "counter_vec"
	Counter      = "counter"
	GaugeVec     = "gauge_vec"
	Gauge        = "gauge"
	HistogramVec = "histogram_vec"
	Histogram    = "histogram"
	SummaryVec   = "summary_vec"
	Summary      = "summary"
)

type Metric struct {
	MetricCollector prometheus.Collector
	Id              string
	Name            string
	Description     string
	Type            string
	Args            []string
}

func NewMetric(subsystem string, m *Metric) prometheus.Collector {
	var metric prometheus.Collector
	switch m.Type {
	case CounterVec:
		metric = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case Counter:
		metric = prometheus.NewCounter(
			prometheus.CounterOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case GaugeVec:
		metric = prometheus.NewGaugeVec(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case Gauge:
		metric = prometheus.NewGauge(
			prometheus.GaugeOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case HistogramVec:
		metric = prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case Histogram:
		metric = prometheus.NewHistogram(
			prometheus.HistogramOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	case SummaryVec:
		metric = prometheus.NewSummaryVec(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
			m.Args,
		)
	case Summary:
		metric = prometheus.NewSummary(
			prometheus.SummaryOpts{
				Subsystem: subsystem,
				Name:      m.Name,
				Help:      m.Description,
			},
		)
	}
	return metric
}

func SetCounterVecMetric(id, name, desc string, args []string) *Metric {
	return &Metric{
		Id:          id,
		Name:        name,
		Description: desc,
		Type:        CounterVec,
		Args:        args,
	}
}

func SetCounterMetric(id, name, desc string) *Metric {
	return &Metric{
		Id:          id,
		Name:        name,
		Description: desc,
		Type:        Counter,
	}
}

func SetGaugeVecMetric(id, name, desc string, args []string) *Metric {
	return &Metric{
		Id:          id,
		Name:        name,
		Description: desc,
		Type:        GaugeVec,
		Args:        args,
	}
}

func SetGaugeMetric(id, name, desc string) *Metric {
	return &Metric{
		Id:          id,
		Name:        name,
		Description: desc,
		Type:        Gauge,
	}
}

func SetHistogramVecMetric(id, name, desc string, args []string) *Metric {
	return &Metric{
		Id:          id,
		Name:        name,
		Description: desc,
		Type:        HistogramVec,
		Args:        args,
	}
}

func SetHistogramMetric(id, name, desc string) *Metric {
	return &Metric{
		Id:          id,
		Name:        name,
		Description: desc,
		Type:        Histogram,
	}
}

func SetSummaryMetric(id, name, desc string) *Metric {
	return &Metric{
		Id:          id,
		Name:        name,
		Description: desc,
		Type:        Summary,
	}
}

func SetSummaryVecMetric(id, name, desc string, args []string) *Metric {
	return &Metric{
		Id:          id,
		Name:        name,
		Description: desc,
		Type:        Summary,
		Args:        args,
	}
}
