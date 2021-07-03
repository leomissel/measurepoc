package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	counterStrategy := MetricStrategy("counter")
	counterMeter := NewMeter(counterStrategy)
	counter := counterMeter.Measure("teste_counter", "TESTE HELPER conter")
	counterImpl := counter.(prometheus.Counter)

	gaugeStrategy := MetricStrategy("gauge")
	gaugeMeter := NewMeter(gaugeStrategy)
	gauge := gaugeMeter.Measure("teste_gauge", "TESTE HELPER gauge")
	gaugeImpl := gauge.(prometheus.Gauge)

	histogramStrategy := MetricStrategy("histogram")
	histogramMeter := NewMeter(histogramStrategy)
	histogram := histogramMeter.Measure("teste_histogram", "TESTE HELPER histogram")
	histogramImpl := histogram.(prometheus.Histogram)

	counterByHelper := HelperNewMesure("counter", "test_by_helper", "HELPER FOR TEST BY HELPER").(prometheus.Counter)

	go func() {
		for {
			time.Sleep(2 * time.Second)
			counterImpl.Inc()
			gaugeImpl.Inc()
			histogramImpl.Observe(prometheus.DefMaxAge.Seconds())
			counterByHelper.Inc()
		}
	}()
}

func main() {

	port := ":2112"
	endpoint := "/metrics"

	recordMetrics()

	log.Println("Server of POC Measeare running at: ", port)
	log.Printf("Endpoint to collector push is http://localhost%s%s\n", port, endpoint)

	if error := runServer(endpoint, port); error != nil {
		log.Fatal("Cannot run http server", error)
	}

}

func runServer(endpoint string, port string) error {
	http.Handle(endpoint, promhttp.Handler())
	return http.ListenAndServe(port, nil)
}
