package main

import (
	"log"
	"os"

	"pharmanet/agent/internal/auto"
	"pharmanet/agent/internal/connectors"
	"pharmanet/agent/internal/connectors/generic_connector"
	"pharmanet/agent/internal/httpserver"
)

func main() {
	token := os.Getenv("AGENT_TOKEN")
	if token == "" {
		token = "CHANGE_ME"
		log.Println("Warning: using default token. Set AGENT_TOKEN env variable for production.")
	}

	// Example: the user selects which apps to give access to
	selectedApps := []string{"generic_pharmacy_app"}

	a := auto.New()
	dets, err := a.Discover(selectedApps)
	if err != nil {
		log.Println("Discovery error:", err)
	}

	if len(dets) == 0 {
		log.Println("No selected software detected. Make sure the recipient has chosen an app.")
	}

	var conns []connectors.Connector
	for _, d := range dets {
		log.Println("Connecting to app:", d.App, "at", d.DSN)
		conns = append(conns, generic_connector.NewGenericConnector(d.App, d.DSN))
	}

	srv := &httpserver.Server{
		Addr:  "127.0.0.1:8081",
		Token: token,
		Search: func(q string, limit int) ([]connectors.StockRecord, error) {
			var agg []connectors.StockRecord
			for _, c := range conns {
				rs, err := c.Search(connectors.SearchRequest{Query: q, Limit: limit})
				if err != nil {
					log.Println("Search error on", c.Name(), ":", err)
					continue
				}
				agg = append(agg, rs...)
			}
			return agg, nil
		},
	}

	log.Println("Starting agent HTTP server on", srv.Addr)
	log.Fatal(srv.Start())
}
