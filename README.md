# sincere_examples
sincere_examples

https://github.com/danclive/sincere

## Branch
RUST: rustc 1.46.0 (04488afe3 2020-08-24)

GO: go version go1.15.2 linux/amd64
cpu: e5-2650 8c
Mem: 8G
SSD
### Sincere
release opt-level = 3
``` 
~$  go-wrk -c 80 -d 5 http://127.0.0.1:8000/hello
Running 5s test @ http://127.0.0.1:8000/hello
  80 goroutine(s) running concurrently
94571 requests in 4.912792899s, 10.01MB read
Requests/sec:		19249.95
Transfer/sec:		2.04MB
Avg Req Time:		4.155855ms
Fastest Request:	132.482µs
Slowest Request:	23.611864ms
Number of Errors:	0
```
#### Gin
``` 
~$ go-wrk -c 80 -d 5 http://127.0.0.1:8081/hello
Running 5s test @ http://127.0.0.1:8081/hello
  80 goroutine(s) running concurrently
181092 requests in 4.904672597s, 19.17MB read
Requests/sec:		36922.34
Transfer/sec:		3.91MB
Avg Req Time:		2.166709ms
Fastest Request:	86.728µs
Slowest Request:	35.10243ms
Number of Errors:	0

```

#### Erguotou github.com/dollarkillerx/erguotou
``` 
~$ go-wrk -c 80 -d 5 http://127.0.0.1:8081/hello
Running 5s test @ http://127.0.0.1:8081/hello
  80 goroutine(s) running concurrently
205570 requests in 4.832811717s, 25.08MB read
Requests/sec:		42536.31
Transfer/sec:		5.19MB
Avg Req Time:		1.880745ms
Fastest Request:	72.469µs
Slowest Request:	32.684381ms
Number of Errors:	0
```
