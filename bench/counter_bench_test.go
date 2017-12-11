// this file exists to document an old attempt at benchmarking counters
package bench

import (
	"testing"

	"github.com/prometheus/client_golang/prometheus"
	metrics "github.com/rcrowley/go-metrics"
)

var fpc prometheus.Counter
var p prometheus.Counter

func init() {
	fpc = prometheus.NewFixedPrecisionCounter(prometheus.Opts{
		Name: "foo",
		Help: "helpful message about foo",
	}, 0)
	p = prometheus.NewCounter(prometheus.CounterOpts{Name: "pfoo", Help: "pfoo help"})
	prometheus.MustRegister(fpc)
	prometheus.MustRegister(p)
}

func BenchmarkGoMetrics(b *testing.B) {
	c := metrics.NewCounter()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			c.Inc(1)
		}
	})
}

func BenchmarkPromCounter(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Inc()
		}
	})
}

func BenchmarkFixedPrecisionCounter(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			fpc.Inc()
		}
	})
}
