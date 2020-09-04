# Traefik Plugin - Bridge

> Work In Progress

The Bridge plugin allows make requests to your custom services. 

You can:

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

