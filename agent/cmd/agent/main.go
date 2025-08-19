package main

import (
	"log"
	"pharmanet/agent/internal/auto"
	"pharmanet/agent/internal/connectors"
	"pharmanet/agent/internal/connectors/generic_connector"
	"pharmanet/agent/internal/httpserver"
)

func main() {
	a := auto.New()
	dets, _ := a.Discover()

	var conns []connectors.Connector
	for _, d := range dets {
		// instantiate a generic connector for any discovered software
		conns = append(conns, generic_connector.NewGenericConnector(d.App, d.DSN))
	}

	srv := &httpserver.Server{
		Addr:  "127.0.0.1:8081",
		Token: "CHANGE_ME",
		Search: func(q string, limit int) ([]connectors.StockRecord, error) {
			var agg []connectors.StockRecord
			for _, c := range conns {
				rs, err := c.Search(connectors.SearchRequest{Query: q, Limit: limit})
				if err == nil {
					agg = append(agg, rs...)
				}
			}
			return agg, nil
		},
	}

	log.Fatal(srv.Start())
}
