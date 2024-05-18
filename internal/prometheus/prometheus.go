package prometheus

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Metrics struct to hold the Prometheus metrics
type Metrics struct {
	RTT        *prometheus.GaugeVec
	PacketLoss *prometheus.GaugeVec
	Jitter     *prometheus.GaugeVec
	HopCount   *prometheus.GaugeVec
	Bandwidth  *prometheus.GaugeVec
}

// NewMetrics creates new Prometheus metrics
func NewMetrics() *Metrics {
	return &Metrics{
		RTT: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "lantern_rtt_seconds",
			Help: "Round-trip time (RTT) in seconds",
		}, []string{"target"}),

		PacketLoss: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "lantern_packet_loss",
			Help: "Packet loss percentage",
		}, []string{"target"}),

		Jitter: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "lantern_jitter_seconds",
			Help: "Jitter in seconds",
		}, []string{"target"}),

		HopCount: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "lantern_hop_count",
			Help: "Hop count",
		}, []string{"target"}),

		Bandwidth: prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "lantern_bandwidth_mbps",
			Help: "Bandwidth in Mbps",
		}, []string{"target", "direction"}),
	}
}

// Register registers the metrics with Prometheus
func (m *Metrics) Register() {
	prometheus.MustRegister(m.RTT)
	prometheus.MustRegister(m.PacketLoss)
	prometheus.MustRegister(m.Jitter)
	prometheus.MustRegister(m.HopCount)
	prometheus.MustRegister(m.Bandwidth)
}

// CollectMetrics is a placeholder function to simulate metrics collection
// In a real implementation, this would collect metrics from other modules
func (m *Metrics) CollectMetrics() {
	// Simulate metric collection
	m.RTT.WithLabelValues("example.com").Set(0.123)
	m.PacketLoss.WithLabelValues("example.com").Set(0.0)
	m.Jitter.WithLabelValues("example.com").Set(0.005)
	m.HopCount.WithLabelValues("example.com").Set(10)
	m.Bandwidth.WithLabelValues("example.com", "download").Set(100)
	m.Bandwidth.WithLabelValues("example.com", "upload").Set(50)
}

// Start initializes and starts the Prometheus HTTP server
func Start() {
	// Set up Prometheus metrics
	metrics := NewMetrics()
	metrics.Register()
	metrics.CollectMetrics()

	// Start Prometheus HTTP server
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Starting Prometheus HTTP server on :8181")
	log.Fatal(http.ListenAndServe(":8181", nil))
}
