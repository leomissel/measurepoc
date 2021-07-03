package main

import "github.com/prometheus/client_golang/prometheus"

type Metric func(string, string) prometheus.Metric

type Meter struct {
	metric Metric
}

func (m Meter) Measure(name string, help string) prometheus.Metric {
	return m.metric(name, help)
}

func NewMeter(metric Metric) Meter {
	return Meter{metric}
}
