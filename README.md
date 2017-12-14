# counters

A small benchmark of existing metrics counters.

## how to run

```bash
% go test -cpu=64 -benchtime=5s -v -bench=. github.com/smcquay/counters/bench
goos: darwin
goarch: amd64
pkg: github.com/smcquay/counters/bench
BenchmarkGoMetrics-64                   500000000               18.3 ns/op
BenchmarkPromCounter-64                 500000000               19.7 ns/op
BenchmarkFixedPrecisionCounter-64       300000000               24.4 ns/op
PASS
ok      github.com/smcquay/counters/bench       32.713s
```

Also when the binary is built:

```bash
$ go install github.com/smcquay/counters
% for c in 1 2 3 4 5 6 7 8 16 32 64 128 256 512 1024 2048; do for i in prom expvar metrics sm; do sudo nice -n -20 ~/bin/counters -conc=$c -dur 20s $i; done; echo; done
prom      :    1 goroutines 7.03E+07 /s got to 1.41E+09 in 20.000082742s
expvar    :    1 goroutines 7.85E+07 /s got to 1.57E+09 in 20.000288233s
metrics   :    1 goroutines 7.31E+07 /s got to 1.46E+09 in 20.000338369s
sm        :    1 goroutines 6.53E+07 /s got to 1.31E+09 in 20.000203838s

prom      :    2 goroutines 2.04E+07 /s got to 4.08E+08 in 20.00034134s
expvar    :    2 goroutines 2.20E+07 /s got to 4.41E+08 in 20.00027187s
metrics   :    2 goroutines 1.93E+07 /s got to 3.86E+08 in 20.00020272s
sm        :    2 goroutines 2.36E+07 /s got to 4.72E+08 in 20.000286071s

prom      :    3 goroutines 2.47E+07 /s got to 4.95E+08 in 20.000428855s
expvar    :    3 goroutines 2.01E+07 /s got to 4.02E+08 in 20.000232194s
metrics   :    3 goroutines 1.74E+07 /s got to 3.47E+08 in 20.000233973s
sm        :    3 goroutines 2.16E+07 /s got to 4.33E+08 in 20.000349826s

prom      :    4 goroutines 2.04E+07 /s got to 4.08E+08 in 20.000160917s
expvar    :    4 goroutines 1.50E+07 /s got to 3.01E+08 in 20.000225511s
metrics   :    4 goroutines 1.56E+07 /s got to 3.12E+08 in 20.000138119s
sm        :    4 goroutines 2.02E+07 /s got to 4.05E+08 in 20.000170586s

prom      :    5 goroutines 2.13E+07 /s got to 4.25E+08 in 20.000307268s
expvar    :    5 goroutines 1.64E+07 /s got to 3.29E+08 in 20.000331322s
metrics   :    5 goroutines 1.58E+07 /s got to 3.15E+08 in 20.000302256s
sm        :    5 goroutines 2.05E+07 /s got to 4.10E+08 in 20.00045397s
 ...
```
