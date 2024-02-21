package main

import (
	"fmt"
	"net/http"
	f "net/url"
	"time"
	"sync"

	"github.com/gorilla/websocket"
)

// HttpClient represents an HTTP client for load testing
type HttpClient struct {
	url string // URL to send HTTP requests
	rps int    // Requests per second
	requestCount int    // Counter for the number of requests sent
    mu           sync.Mutex
}

// Start starts the HTTP client
func (c *HttpClient) Start() {
	// Check if requests per second is set
	if c.rps == 0 {
		fmt.Println("No requests per second selected")
		return
	}
	// Create a ticker to regulate request frequency
	ticker := time.NewTicker(time.Second / time.Duration(c.rps))
	go func() {
		for range ticker.C {
			c.makeRequest()
		}
	}()
}

// makeRequest sends an HTTP GET request to the specified URL
func (c *HttpClient) makeRequest() {
	resp, err := http.Get(c.url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()
	c.mu.Lock()
	c.requestCount++ // Increment request count
	c.mu.Unlock()
	fmt.Println("HTTP request sent. Total requests:", c.requestCount)
}

// WebsocketClient represents a WebSocket client for load testing
type WebsocketClient struct {
	url  string          // WebSocket URL
	rps  int             // Requests per second
	conn *websocket.Conn // WebSocket connection
	mu   sync.Mutex
}

// Start starts the WebSocket client
func (c *WebsocketClient) Start() {
	var err error
	// Establish WebSocket connection
	c.conn, _, err = websocket.DefaultDialer.Dial(c.url, nil)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	// Create a ticker to regulate request frequency
	ticker := time.NewTicker(time.Second / time.Duration(c.rps))
	go func() {
		for range ticker.C {
			c.requestUpdate()
		}
	}()
}

// requestUpdate sends a ping message over the WebSocket connection
func (c *WebsocketClient) requestUpdate() {
	c.mu.Lock()
	err := c.conn.WriteMessage(websocket.TextMessage, []byte("ping"))
	c.mu.Unlock()
	if  err != nil {
		fmt.Println("Error:", err)
	}
}

// LoadTest simulates load testing on the provided URL with the given concurrency and requests per second
func LoadTest(url string, concurrency, rps int) {

	u, err := f.Parse(url)
	if err != nil {
		panic(err)
	}
 // Start multiple clients concurrently
 for i := 0; i < concurrency; i++ {
	var client interface {
		Start()
	}
	// Determine the type of client based on the URL scheme
	if u.Scheme == "ws" || u.Scheme == "wss" {
		client = &WebsocketClient{url: url, rps: rps}
	} else {
		client = &HttpClient{url: url, rps: rps}
	}
	go client.Start()
	time.Sleep(100 * time.Millisecond) // Add a short delay between starting clients
}
}
