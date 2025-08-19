package connectors

import "fmt"

// GenericConnector is a placeholder template for any pharmacy software
type GenericConnector struct {
	NameStr string
	DSN     string
}

func NewGenericConnector(name, dsn string) *GenericConnector {
	return &GenericConnector{NameStr: name, DSN: dsn}
}

func (c *GenericConnector) Name() string {
	return c.NameStr
}

func (c *GenericConnector) Search(req SearchRequest) ([]StockRecord, error) {
	// TODO: Implement the real connection logic based on DSN and software type
	fmt.Println("Searching", req.Query, "on", c.NameStr)
	return []StockRecord{
		{ProductID: "123", ProductName: req.Query, Quantity: 50, LastUpdated: "2025-08-19T00:00:00Z"},
	}, nil
}
