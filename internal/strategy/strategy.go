package strategy

import "fmt"

// BaseStrategy defines the interface for any connector
type BaseStrategy interface {
    Connect() error
    SearchInventory(query string) ([]string, error)
}

// Example placeholder strategy
type GenericStrategy struct {
    AppName string
}

func (g *GenericStrategy) Connect() error {
    fmt.Println("Connecting to app:", g.AppName)
    // TODO: dynamically pick strategy (SQL, API, UI automation)
    return nil
}

func (g *GenericStrategy) SearchInventory(query string) ([]string, error) {
    fmt.Println("Searching inventory for:", query)
    // TODO: implement search logic
    return []string{"ExampleProduct1", "ExampleProduct2"}, nil
}
