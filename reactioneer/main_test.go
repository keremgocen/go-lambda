package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestHandler(t *testing.T) {
	// t.Run("Unable to get IP", func(t *testing.T) {
	// 	DefaultHTTPGetAddress = "http://127.0.0.1:12345"

	// 	_, err := handler(events.APIGatewayProxyRequest{})
	// 	if err == nil {
	// 		t.Fatal("Error failed to trigger with an invalid request")
	// 	}
	// })

	// t.Run("Non 200 Response", func(t *testing.T) {
	// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(500)
	// 	}))
	// 	defer ts.Close()

	// 	DefaultHTTPGetAddress = ts.URL

	// 	_, err := handler(events.APIGatewayProxyRequest{})
	// 	if err != nil && err.Error() != ErrNon200Response.Error() {
	// 		t.Fatalf("Error failed to trigger with an invalid HTTP response: %v", err)
	// 	}
	// })

	// t.Run("Unable decode IP", func(t *testing.T) {
	// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		w.WriteHeader(500)
	// 	}))
	// 	defer ts.Close()

	// 	DefaultHTTPGetAddress = ts.URL

	// 	_, err := handler(events.APIGatewayProxyRequest{})
	// 	if err == nil {
	// 		t.Fatal("Error failed to trigger with an invalid HTTP response")
	// 	}
	// })

	t.Run("Successful Request", func(t *testing.T) {
		os.Setenv("SLACK_TOKEN", "VALUE")

		ts := httptest.NewServer(
			http.HandlerFunc(
				func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(http.StatusBadGateway)
					w.Header().Set("Content-Type", "application/json")
				}))
		defer ts.Close()

		request := SlackRequest{
			Token:     "VALUE",
			EventType: "alo",
		}

		event, err := handler(request)
		fmt.Println(event.StatusCode)
		fmt.Println(event.Headers)
		fmt.Println(event.Body)
		if err != nil {
			t.Fatal("Everything should be ok")
		}
	})
}
