package main

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var CounterImpl Metric = Metric(func(name string, help string) prometheus.Metric {
	return promauto.NewCounter(prometheus.CounterOpts{
		Name: name,
		Help: help,
	})
})

var gaugeImpl Metric = Metric(func(name string, help string) prometheus.Metric {
	return promauto.NewGauge(prometheus.GaugeOpts{
		Name: name,
		Help: help,
	})
})

var histImpl Metric = Metric(func(name string, help string) prometheus.Metric {
	return promauto.NewHistogram(prometheus.HistogramOpts{
		Name: name,
		Help: help,
	})
})

func MetricStrategy(strategy string) Metric {
	var metric Metric
	switch strategy {
	case "counter":
		metric = CounterImpl
	case "gauge":
		metric = gaugeImpl
	case "histogram":
		metric = histImpl
	default:
		log.Fatal("Please, choose a type of your metric: counter, gauge or histogram")
	}

	return metric
}

func HelperNewMesure(metric string, name string, helper string) prometheus.Metric {
	strategy := MetricStrategy(metric)
	meter := NewMeter(strategy)
	mesuarement := meter.Measure(name, helper)
	return mesuarement
}
