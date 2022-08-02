package exporter

import (
	"github.com/prometheus/client_golang/prometheus"
)

type gpuCollector struct {
	gpuTop *prometheus.Desc
}

func newGpuCollector() *gpuCollector {
	return &gpuCollector{
		gpuTop: prometheus.NewDesc(prometheus.BuildFQName("gpu", "top", "top"),
			"GPU Utility Rate",
			[]string{"top"}, nil,
		),
	}
}

func (collector *gpuCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.gpuTop
}

func (collector *gpuCollector) Collect(ch chan<- prometheus.Metric) {

	var metricValue float64
	if 1 == 1 {
		metricValue = 0.5
	}

	ch <- prometheus.MustNewConstMetric(collector.gpuTop, prometheus.GaugeValue, metricValue, "top")
}

func Register() {
	collector := newGpuCollector()
	prometheus.MustRegister(collector)
}
