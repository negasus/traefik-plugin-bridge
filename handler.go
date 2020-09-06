package traefik_plugin_bridge

import (
	"fmt"
	"net/http"
)

func (bridge *Bridge) ServeHTTP(rw http.ResponseWriter, httpReq *http.Request) {
	req := AcquireRequest()
	defer ReleaseRequest(req)
	req.FillFromHTTPRequest(httpReq, &bridge.requestConfig)

	var resp *Response
	var err error

	switch bridge.mode {
	case bridgeModeBINARY:
		resp, err = bridge.clientBINARY.Call(req)
	default:
		resp, err = bridge.clientHTTP.Call(req)
	}
	if err != nil {
		bridge.log("error call backend, %s", err)
		bridge.next.ServeHTTP(rw, httpReq)
		return
	}
	defer releaseResponse(resp)

	for _, value := range resp.DeleteRequestHeaders {
		httpReq.Header.Del(value)
	}
	for name, value := range resp.AddRequestHeaders {
		httpReq.Header.Add(name, value)
	}
	for name, value := range resp.AddResponseHeaders {
		rw.Header().Add(name, value)
	}
	if resp.InterruptRequest != nil {
		rw.WriteHeader(resp.InterruptRequest.ResponseCode)
		if len(resp.InterruptRequest.Body) > 0 {
			n, err := rw.Write([]byte(resp.InterruptRequest.Body))
			if err != nil {
				bridge.log("error write body, %s", err)
			}
			if n != len(resp.InterruptRequest.Body) {
				bridge.log("write unexpected body length", fmt.Errorf("write %d, expect %d", n, len(resp.InterruptRequest.Body)))
			}
		}
		return
	}

	bridge.next.ServeHTTP(rw, httpReq)
}
