// internal/strategy/factory.go (add cache support)
type GenericStrategy struct {
    AppName string
    Cache   *cache.Cache
}

func (g *GenericStrategy) Connect() error {
    fmt.Println("Connecting generically to:", g.AppName)
    c, _ := cache.NewCache("./agent_cache.db")
    g.Cache = c
    return nil
}

func (g *GenericStrategy) SearchInventory(query string) ([]string, error) {
    // Simulate live search
    results := []string{"Product1", "Product2"} // Replace with real logic
    // Save to cache
    g.Cache.Save(query, results)
    return results, nil
}
