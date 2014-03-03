gobitrand
=========

Consume parts of the entropy bits at a time.

## Install
1. Make sure your GOPATH environment variable is set
2. go get github.com/PhillipNordwal/gobitrand
3. go install github.com/PhillipNordwal/gobitrand

## Test
1. go test github.com/PhillipNordwal/gobitrand

## Benchmark
1. go test github.com/PhillipNordwal/gobitrand -bench=".*"

## Profile
1. go test github.com/PhillipNordwal/gobitrand -bench="BenchmarkTwo_bits_speed" -cpuprofile Two_bit.cpu.out
2. go tool pprof $GOPATH/src/github.com/PhillipNordwal/gobitrand/gobitrand.test $GOPATH/src/github.com/PhillipNordwal/gobitrand/Two_bit.cpu.out
3. top10
4. top10 -cum
5. web
6. weblist Two_bit
