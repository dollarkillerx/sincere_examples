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
~$ go-wrk -c 80 -d 5 http://127.0.0.1:8001/hello
Running 5s test @ http://127.0.0.1:8081/hello
  80 goroutine(s) running concurrently
303182 requests in 4.936371019s, 32.09MB read
Requests/sec:		61417.99
Transfer/sec:		6.50MB
Avg Req Time:		1.302549ms
Fastest Request:	32.433µs
Slowest Request:	46.306059ms
Number of Errors:	0
```

#### Erguotou github.com/dollarkillerx/erguotou
``` 
~$ go-wrk -c 80 -d 5 http://127.0.0.1:8001/hello
Running 5s test @ http://127.0.0.1:8081/hello
  80 goroutine(s) running concurrently
419923 requests in 4.895925827s, 51.63MB read
Requests/sec:		85769.89
Transfer/sec:		10.54MB
Avg Req Time:		932.728µs
Fastest Request:	25.092µs
Slowest Request:	29.157395ms
Number of Errors:	0
```
