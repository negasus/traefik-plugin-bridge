# Traefik Plugin - Bridge

> Work In Progress

> `BINARY` mode still not implemented

Write middlewares with you preferred languages!

PHP, Python, NodeJS or more - it does not matter!
 
The Bridge plugin allows make requests to your services.

## You can

- get the request data
    - headers
    - remote address
    - method
    - uri
    
- modify the request
    - add request headers
    - remote request headers
    - add response headers
    - interrupt the request with custom body and status code 

## Usage example

1. Run `logserver` ([github](https://github.com/negasus/logserver)) as backend for logging requests (or any you prefer software)

```
docker run --rm -p 2000:2000 negasus/logserver
```

2. Run our service for listen the middleware and modify the request

```
go run ./test/http
```

3. Run the Traefik with a plugin configuration:

```
address: 'http://127.0.0.1:3000'
```

4. Send request to the Traefik

```
# in my case the Traefik is listen 8000 port

curl http://localhost:8000 
```

5. Check `logserver` output

```
----------[ 1 ]----------
2020-09-06 17:33:29.3858593 +0000 UTC m=+172.295935301
[172.17.0.1:47380] GET /

Host: localhost:8000
Content-Length: 0
User-Agent: Middleware!             <----- It's work!
Accept: */*
X-Forwarded-For: 127.0.0.1
X-Forwarded-Host: localhost:8000
X-Forwarded-Port: 8000
X-Forwarded-Proto: http
X-Forwarded-Server: notebook
X-Real-Ip: 127.0.0.1
Accept-Encoding: gzip
```

## Configuration

```
protocol: 'HTTP' or 'BINARY'. Default 'HTTP'
address: Address of your service. Default 'http://127.0.0.1:8000'
timeout: Request timeout to your service, ms. Default '500'
fail_strategy: Strategy on fail request to your service. Default 'PASS'
request:
    - headers: bool. Send request headers to your service. Default 'true'
    - remote_address: bool. Send remote address to your service. Default 'true'
    - method: bool. Send http method to your service. Default 'true'
    - uri: bool. Send URI method to your service. Default 'true'
```

## Request

```
type Request struct {
	Headers    map[string][]string `json:"1,omitempty"`
	RemoteAddr string              `json:"2,omitempty"`
	Method     string              `json:"3,omitempty"`
	RequestURI string              `json:"4,omitempty"`
}
```


## Response

```
type Response struct {
	AddRequestHeaders    map[string]string         `json:"1,omitempty"`
	AddResponseHeaders   map[string]string         `json:"2,omitempty"`
	DeleteRequestHeaders []string                  `json:"3,omitempty"`
	InterruptRequest     *ResponseInterruptRequest `json:"4,omitempty"`
}

type ResponseInterruptRequest struct {
	ResponseCode int    `json:"1,omitempty"`
	Body         string `json:"2,omitempty"`
}
```

## Probable Roadmap

- implementation for the `BINARY` mode
- full documentation 
- check the response status from the customer service
- multiple connections to services
- circuit breaker
