# sincere_examples
sincere_examples

https://github.com/danclive/sincere

## Branch
RUST: rustc 1.46.0 (04488afe3 2020-08-24)

GO: go version go1.15.2 linux/amd64
cpu: e5-2650 8c
Mem: 8G
SSD
#### Sincere
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

### 测试环境2 实体机
i3-9100
8G
HHD
go version go1.13.5 linux/amd64
rustc 1.46.0 (04488afe3 2020-08-24)
#### Sincere
release opt-level = 3
``` 
Running 5s test @ http://127.0.0.1:8000/hello
  80 goroutine(s) running concurrently
290245 requests in 4.938527955s, 30.72MB read
Requests/sec:		58771.56
Transfer/sec:		6.22MB
Avg Req Time:		1.361202ms
Fastest Request:	41.236µs
Slowest Request:	31.390056ms
Number of Errors:	0
```

#### Gin
``` 
Running 5s test @ http://127.0.0.1:8081/hello
  80 goroutine(s) running concurrently
310550 requests in 4.93030957s, 32.87MB read
Requests/sec:		62987.93
Transfer/sec:		6.67MB
Avg Req Time:		1.270084ms
Fastest Request:	28.931µs
Slowest Request:	50.370903ms
Number of Errors:	0
```

#### Erguotou
``` 
Running 5s test @ http://127.0.0.1:8081/hello
  80 goroutine(s) running concurrently
441697 requests in 4.904608774s, 54.31MB read
Requests/sec:		90057.54
Transfer/sec:		11.07MB
Avg Req Time:		888.32µs
Fastest Request:	25.2µs
Slowest Request:	26.29641ms
Number of Errors:	0
```