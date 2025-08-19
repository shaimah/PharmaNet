package connectors

// Request to search inventory
type SearchRequest struct {
	Query string
	Limit int
}

// Unified stock record
type StockRecord struct {
	ProductID, ProductName, Strength, PackSize, Location string
	Quantity                                             int
	ExpiryDate, LastUpdated                               string
}

// Connector interface — every pharmacy software implements this
type Connector interface {
	Name() string
	Search(SearchRequest) ([]StockRecord, error)
}

// Detection result from auto-discovery
type Detection struct {
	App  string            // e.g., "pharmacy_software_name"
	Kind string            // "sqlite" | "postgres" | "api"
	DSN  string            // path or connection string
	Meta map[string]string // extra info
}

// Detector interface — finds installed apps on this machine
type Detector interface {
	App() string
	Detect() ([]Detection, error)
}
