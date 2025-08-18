package httpserver

import (
    "fmt"
    "net/http"
    "pharmanet/internal/strategy"
)

var strat strategy.BaseStrategy

func StartServer(port int, s strategy.BaseStrategy) {
    strat = s
    http.HandleFunc("/v1/inventory/search", searchHandler)
    addr := fmt.Sprintf(":%d", port)
    fmt.Println("HTTP API listening on", addr)
    http.ListenAndServe(addr, nil)
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    results, err := strat.SearchInventory(query)
    if err != nil {
        // fallback to cache if live search fails
        if gs, ok := strat.(*strategy.GenericStrategy); ok {
            results, _ = gs.Cache.Search(query)
        }
    }
    fmt.Fprintf(w, "%v", results)
}
