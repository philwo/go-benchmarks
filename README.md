## Results

### slices_test

#### go 1.32.2 on Apple M2 Ultra
```
$ go test . && go test -bench .
ok  	github.com/philwo/go-benchmarks	0.019s
goos: darwin
goarch: arm64
pkg: github.com/philwo/go-benchmarks
cpu: Apple M2 Ultra
BenchmarkIndexBytesAny-24                    	50747736	        23.36 ns/op
BenchmarkIndexBytesAnyIndexFuncDirect-24     	12743683	        93.54 ns/op
BenchmarkIndexBytesAnyIndexFuncWrapped-24    	54417345	        22.21 ns/op
BenchmarkSkipBytesAny-24                     	51949456	        23.30 ns/op
BenchmarkSkipBytesAnyIndexFunc-24            	54437094	        22.03 ns/op
PASS
ok  	github.com/philwo/go-benchmarks	6.203s
```
