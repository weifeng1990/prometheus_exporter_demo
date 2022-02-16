package collector
import (
	"github.com/prometheus/client_golang/prometheus"
)

type collector struct {
	g *prometheus.Desc
}

func Init() *collector{
	return &collector{
		g: prometheus.NewDesc(
			prometheus.BuildFQName("test", "hello", "world"),
			"test hello world",
			[]string{"test"}, nil,
		),
	}
}

func (c collector) Describe(ch chan<- *prometheus.Desc) {
	ch <- prometheus.NewDesc("test", "test", nil, nil)
}

func (c collector) Collect(ch chan<- prometheus.Metric){
	ch <- prometheus.MustNewConstMetric(
		c.g, prometheus.GaugeValue, float64(32), "wei")
}
