package main
import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/version"
	"github.com/prometheus/common/promlog"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"test/collector"
	"os"
	"net/http"
	"github.com/prometheus/exporter-toolkit/web"
)

func handler(w http.ResponseWriter, r *http.Request){
	registry := prometheus.NewRegistry()
	c := collector.Init()
	registry.MustRegister(c)
	h := promhttp.HandlerFor(registry, promhttp.HandlerOpts{})
	h.ServeHTTP(w, r)	
}

func main(){
	promlogConfig := &promlog.Config{}
	logger := promlog.New(promlogConfig)
	prometheus.MustRegister(version.NewCollector("test_exporter"))
	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		handler(w, r)
	})
		srv := &http.Server{Addr: ":8080"}
	if err := web.ListenAndServe(srv, "", logger); err != nil {
		os.Exit(1)
	}
}
