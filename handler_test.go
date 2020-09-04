package traefik_plugin_bridge

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testserver struct {
	called bool

	requestHeaders http.Header
	requestBody    []byte

	responseHeaders http.Header
	responseBody    []byte
	responseCode    int
}

func (t *testserver) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	t.called = true
	t.requestHeaders = req.Header
	t.requestBody, _ = ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	for key, value := range t.responseHeaders {
		rw.Header().Add(key, value[0])
	}
	rw.WriteHeader(t.responseCode)
	rw.Write(t.responseBody)
}

func TestHandler(t *testing.T) {
	backendServer := &testserver{
		responseHeaders: map[string][]string{"RESP1": []string{"1"}},
	}

	responseFromService := &Response{
		AddRequestHeaders:    map[string]string{"NEW_REQ_HEADER": "3"},
		AddResponseHeaders:   map[string]string{"NEW_RESP_HEADER": "3"},
		DeleteRequestHeaders: []string{"HEADER1"},
		InterruptRequest:     nil,
	}
	responseFromServiceBytes, _ := responseFromService.marshalJSON()

	serviceServer := &testserver{
		responseCode: 200,
		responseBody: responseFromServiceBytes,
	}
	server := httptest.NewServer(serviceServer)
	defer server.Close()

	b := &Bridge{
		clientHTTP: &ClientHTTP{
			client:  &http.Client{},
			address: server.URL,
		},
		next: backendServer,
	}

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "http://domain.com", nil)
	req.Header.Add("HEADER1", "1")
	req.Header.Add("HEADER2", "2")

	b.ServeHTTP(rw, req)

	if rw.Header().Get("RESP1") != "1" {
		t.Fatalf("unexpected value")
	}
	if rw.Header().Get("NEW_RESP_HEADER") != "3" {
		t.Fatalf("unexpected value")
	}

	if backendServer.requestHeaders.Get("HEADER1") != "" {
		t.Fatalf("unexpected value")
	}
	if backendServer.requestHeaders.Get("HEADER2") != "2" {
		t.Fatalf("unexpected value")
	}
	if backendServer.requestHeaders.Get("NEW_REQ_HEADER") != "3" {
		t.Fatalf("unexpected value")
	}
}

func TestHandler_Interrupt(t *testing.T) {
	backendServer := &testserver{}

	responseFromService := &Response{
		InterruptRequest: &ResponseInterruptRequest{
			ResponseCode: 403,
			Body:         "forbidden",
		},
	}
	responseFromServiceBytes, _ := responseFromService.marshalJSON()

	serviceServer := &testserver{
		responseCode: 200,
		responseBody: responseFromServiceBytes,
	}
	server := httptest.NewServer(serviceServer)
	defer server.Close()

	b := &Bridge{
		clientHTTP: &ClientHTTP{
			client:  &http.Client{},
			address: server.URL,
		},
		next: backendServer,
	}

	rw := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "http://domain.com", nil)

	b.ServeHTTP(rw, req)

	if rw.Code != 403 {
		t.Fatalf("unexpected value")
	}
	if rw.Body.String() != "forbidden" {
		t.Fatalf("unexpected value")
	}
	if backendServer.called {
		t.Fatalf("unexpected value")
	}
}
