## Results

### slices_test

#### go 1.32.2 on Apple M2 Ultra (macOS)
```
ok  	github.com/philwo/go-benchmarks	0.024s
goos: darwin
goarch: arm64
pkg: github.com/philwo/go-benchmarks
cpu: Apple M2 Ultra
BenchmarkIndexBytesAny-24                                 	51002155	        23.24 ns/op
BenchmarkIndexBytesAnyDoesntContain-24                    	50088610	        24.02 ns/op
BenchmarkIndexBytesAnyImmediateHit-24                     	1000000000	         1.197 ns/op
BenchmarkIndexBytesAnyIndexFuncDirect-24                  	13096927	        91.40 ns/op
BenchmarkIndexBytesAnyIndexFuncDirectDoesntContain-24     	12655518	        94.36 ns/op
BenchmarkIndexBytesAnyIndexFuncDirectImmediateHit-24      	490879358	         2.458 ns/op
BenchmarkIndexBytesAnyIndexFuncWrapped-24                 	54509527	        22.04 ns/op
BenchmarkIndexBytesAnyIndexFuncWrappedDoesntContain-24    	51553311	        22.80 ns/op
BenchmarkIndexBytesAnyIndexFuncWrappedImmediateHit-24     	625441933	         1.914 ns/op
PASS
ok  	github.com/philwo/go-benchmarks	11.617s
```

#### go 1.32.2 on Threadripper Pro 3995WX (Linux)
```
ok  	github.com/philwo/go-benchmarks	0.013s
goos: linux
goarch: amd64
pkg: github.com/philwo/go-benchmarks
cpu: AMD Ryzen Threadripper PRO 3995WX 64-Cores     
BenchmarkIndexBytesAny-128                                  	47714817	        25.27 ns/op
BenchmarkIndexBytesAnyDoesntContain-128                     	44515806	        27.26 ns/op
BenchmarkIndexBytesAnyImmediateHit-128                      	1000000000	         1.025 ns/op
BenchmarkIndexBytesAnyIndexFuncDirect-128                   	15786864	        75.32 ns/op
BenchmarkIndexBytesAnyIndexFuncDirectDoesntContain-128      	15742936	        76.07 ns/op
BenchmarkIndexBytesAnyIndexFuncDirectImmediateHit-128       	348945354	         3.306 ns/op
BenchmarkIndexBytesAnyIndexFuncWrapped-128                  	51606722	        22.93 ns/op
BenchmarkIndexBytesAnyIndexFuncWrappedDoesntContain-128     	49247053	        24.29 ns/op
BenchmarkIndexBytesAnyIndexFuncWrappedImmediateHit-128      	496932770	         2.267 ns/op
PASS
ok  	github.com/philwo/go-benchmarks	15.527s
```
