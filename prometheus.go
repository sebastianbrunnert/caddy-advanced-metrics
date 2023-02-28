package advanced_metrics

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func StartServer(port int) {
	if _, ok := counters[port]; !ok {
		counters[port] = prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Name: "caddy_advanced_metrics_requests",
				Help: "Number of requests",
			},
			[]string{"method", "path", "status", "host"},
		)
		reg := prometheus.NewRegistry()
		reg.MustRegister(counters[port])
		http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))
		go http.ListenAndServe(":"+strconv.Itoa(port), nil)
	}
}

func HandleRequest(port int, method string, path string, status int, host string) {
	fmt.Println("HandleRequest", port, method, path, status, host)
	if _, ok := counters[port]; ok {
		counters[port].WithLabelValues(method, path, strconv.Itoa(status), host).Inc()
	}
}

var (
	counters = make(map[int]*prometheus.CounterVec)
)
