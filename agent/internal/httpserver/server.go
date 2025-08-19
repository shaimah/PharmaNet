package httpserver

import (
    "fmt"
    "net/http"
)

// StartServer runs the local HTTP API
func StartServer(port int) {
    http.HandleFunc("/v1/inventory/search", func(w http.ResponseWriter, r *http.Request) {
        query := r.URL.Query().Get("q")
        fmt.Fprintf(w, "Search results for: %s\n", query)
    })

    addr := fmt.Sprintf("127.0.0.1:%d", port)
    fmt.Println("HTTP server running on", addr)
    http.ListenAndServe(addr, nil)
}
