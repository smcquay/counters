package main

import (
	"expvar"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	dto "github.com/prometheus/client_model/go"
	metrics "github.com/rcrowley/go-metrics"
	"github.com/smcquay/prom"
)

const usage = "fpc <expvar|prom|metrics|sm>"

var duration = flag.Duration("dur", 10*time.Second, "how long to run test")
var conc = flag.Int("conc", runtime.NumCPU()-1, "how many goroutines to launch")
var sched = flag.Int("sched", 10000, "number of increments to perform before forcing go scheduler yield")

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		fmt.Fprintf(os.Stderr, "%v\n", usage)
		os.Exit(1)
	}
	test := flag.Arg(0)

	concurrency := *conc

	var printFunc func(dur time.Duration)

	switch test {
	case "expvar":
		counter := expvar.NewInt("test")
		printFunc = func(dur time.Duration) {
			render(test, *conc, float64(counter.Value()), dur)
		}
		for i := 0; i < concurrency; i++ {
			go func() {
				for j := 0; ; j++ {
					if j%*sched == 0 {
						runtime.Gosched()
					}
					counter.Add(1)
				}
			}()
		}
	case "metrics":
		c := metrics.NewCounter()
		printFunc = func(dur time.Duration) {
			render(test, *conc, float64(c.Count()), dur)
		}
		for i := 0; i < concurrency; i++ {
			go func() {
				for j := 0; ; j++ {
					if j%*sched == 0 {
						runtime.Gosched()
					}
					c.Inc(1)
				}
			}()
		}
	case "prom":
		p := prometheus.NewCounter(
			prometheus.CounterOpts{
				Name: "counter",
				Help: "A counter metric",
			},
		)
		printFunc = func(dur time.Duration) {
			metric := &dto.Metric{}
			p.Write(metric)
			render(test, *conc, *metric.Counter.Value, dur)
		}
		for i := 0; i < concurrency; i++ {
			go func() {
				for j := 0; ; j++ {
					if j%*sched == 0 {
						runtime.Gosched()
					}
					p.Inc()
				}
			}()
		}
	case "sm":
		p := prom.NewCounter(
			prometheus.CounterOpts{
				Name: "counter",
				Help: "A counter metric",
			},
			2,
		)
		printFunc = func(dur time.Duration) {
			metric := &dto.Metric{}
			p.Write(metric)
			render(test, *conc, *metric.Counter.Value, dur)
		}
		for i := 0; i < concurrency; i++ {
			go func() {
				for j := 0; ; j++ {
					if j%*sched == 0 {
						runtime.Gosched()
					}
					p.Add(1)
				}
			}()
		}
	default:
		fmt.Fprintf(os.Stderr, "unknown test\n")
		os.Exit(1)
	}

	s := time.Now()
	time.Sleep(*duration)
	printFunc(time.Since(s))
}

func render(name string, conc int, val float64, dur time.Duration) {
	aps := val / dur.Seconds()
	fmt.Printf("%-10s: %4d goroutines %0.2E /s got to %0.2E in %s\n", name, conc, aps, val, dur)
}
