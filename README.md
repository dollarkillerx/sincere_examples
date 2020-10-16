# sincere_examples
sincere_examples

https://github.com/danclive/sincere

## Branch

### Sincere
``` 
~$  go-wrk -c 80 -d 5 http://127.0.0.1:8000/hello
Running 5s test @ http://127.0.0.1:8000/hello
  80 goroutine(s) running concurrently
94569 requests in 4.995693174s, 10.01MB read
Requests/sec:		18930.11
Transfer/sec:		2.00MB
Avg Req Time:		4.226072ms
Fastest Request:	161.973µs
Slowest Request:	40.050931ms
Number of Errors:	0
```
#### Gin
``` 
~$ go-wrk -c 80 -d 5 http://127.0.0.1:8000/hello
Running 5s test @ http://127.0.0.1:8000/hello
  80 goroutine(s) running concurrently
96390 requests in 4.997789904s, 10.20MB read
Requests/sec:		19286.53
Transfer/sec:		2.04MB
Avg Req Time:		4.147973ms
Fastest Request:	189.76µs
Slowest Request:	19.825361ms
Number of Errors:	0
```

#### Erguotou github.com/dollarkillerx/erguotou
``` 
~$ go-wrk -c 80 -d 5 http://127.0.0.1:8000/hello
Running 5s test @ http://127.0.0.1:8000/hello
  80 goroutine(s) running concurrently
94569 requests in 4.995693174s, 10.01MB read
Requests/sec:		18930.11
Transfer/sec:		2.00MB
Avg Req Time:		4.226072ms
Fastest Request:	161.973µs
Slowest Request:	40.050931ms
Number of Errors:	0
```