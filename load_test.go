package main

import (
    "net/http"
    "net/http/httptest"
    "testing"
    "time"

    "github.com/gorilla/websocket"
)

// type MockHttpClient struct {
//     url string
//     rps int
// }

// func (c *MockHttpClient) Start() {}

// type MockWebsocketClient struct {
//     url string
//     rps int
// }

// func (c *MockWebsocketClient) Start() {}

func TestLoadTest(t *testing.T) {
    // Mock HTTP server
    httpServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
    }))
    defer httpServer.Close()

    // Mock WebSocket server
    wsServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        upgrader := websocket.Upgrader{}
        _, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            t.Fatal(err)
        }
    }))
    defer wsServer.Close()

    tests := []struct {
        name        string
        url         string
        concurrency int
        rps         int
    }{
        {
            name:        "HTTP Load Test",
            url:         httpServer.URL,
            concurrency: 5,
            rps:         10,
        },
        {
            name:        "WebSocket Load Test",
            url:         "ws" + wsServer.URL[4:], // Convert http:// to ws://
            concurrency: 5,
            rps:         10,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            LoadTest(tt.url, tt.concurrency, tt.rps)
            time.Sleep(1 * time.Second) // Allow some time for the goroutines to start
        })
    }
}