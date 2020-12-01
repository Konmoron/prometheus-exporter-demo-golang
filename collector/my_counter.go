package collector

import (
	"math/rand"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// https://prometheus.io/docs/guides/go-application/#instrumenting-a-go-application-for-prometheus

// RecordMetrics add
func RecordMetrics() {
	go func() {
		for {
			OpsProcessed.Inc()
			time.Sleep(time.Duration(rand.Int63n(1000)) * time.Millisecond)
		}
	}()
}

var (
	// OpsProcessed OpsProcessed
	OpsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name:        "myapp_processed_ops_total",
		Help:        "The total number of processed events",
		ConstLabels: map[string]string{"test_node_exporter_label": "test_node_exporter_value", "app_name": "test_node_exporter"},
	})
)

// func main() {
//         recordMetrics()

//         http.Handle("/metrics", promhttp.Handler())
//         http.ListenAndServe(":2112", nil)
// }
