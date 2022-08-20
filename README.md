# go-sha256-kernel

A library like `crypto/sha256` using Linux kernel crypto API.

## Benchmark

`crpyto/sha256` is faster than this library :)

```
$ go test --bench . --benchmem 
goos: linux
goarch: amd64
pkg: github.com/nokute78/go-sha256-kernel
cpu: AMD Ryzen 7 PRO 3700U w/ Radeon Vega Mobile Gfx
BenchmarkGoSum256     	  689259	      1692 ns/op	       0 B/op	       0 allocs/op
BenchmarkKernelSum256 	   19651	     58921 ns/op	     656 B/op	      14 allocs/op
PASS
ok  	github.com/nokute78/go-sha256-kernel	2.970s
```

## Reference

- Kernel Doc
  - https://www.kernel.org/doc/html/latest/crypto/userspace-if.html
- SockaddrALG usage
  - https://pkg.go.dev/golang.org/x/sys/unix#SockaddrALG

## License

BSD 3-Clause "New" or "Revised" License
