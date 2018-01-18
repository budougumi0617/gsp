# 8.4 UnixドメインソケットとTCPのベンチマーク
UnixドメインソケットとTCPソケットのベンチマークを比較する。

# Result
```bash
$ go test -bench .
goos: darwin
goarch: amd64
pkg: github.com/budougumi0617/gsp/ch08/benchmark
BenchmarkTCPSercer-4         	    1000	   7547749 ns/op
BenchmarkUDSStreamServer-4   	    5000	    246761 ns/op
PASS
ok  	github.com/budougumi0617/gsp/ch08/benchmark	9.968s

$ system_profiler SPHardwareDataType
Hardware:

    Hardware Overview:

      Model Name: MacBook Pro
      Model Identifier: MacBookPro12,1
      Processor Name: Intel Core i5
      Processor Speed: 2.9 GHz
      Number of Processors: 1
      Total Number of Cores: 2
      L2 Cache (per Core): 256 KB
      L3 Cache: 3 MB
      Memory: 16 GB
      Boot ROM Version: M
      SMC Version (system): 2.28f7
      Serial Number (system): 
      Hardware UUID: 
```
