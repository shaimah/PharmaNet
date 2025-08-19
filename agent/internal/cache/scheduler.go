package cache

import (
	"time"
	"pharmanet/agent/internal/connectors"
	"pharmanet/agent/internal/push"
)

type Scheduler struct {
	Conns  []connectors.Connector
	Pusher *push.Client
	Store  *Store
}

func (s *Scheduler) Start() {
	ticker := time.NewTicker(30 * time.Minute)
	for range ticker.C {
		var all []connectors.StockRecord
		for _, c := range s.Conns {
			if rs, err := c.Search(connectors.SearchRequest{Query: "", Limit: 5000}); err == nil {
				all = append(all, rs...)
			}
		}
		_ = s.Store.Save(all)        // local cache
		_ = s.Pusher.PushSnapshot(all) // push to central server
	}
}
