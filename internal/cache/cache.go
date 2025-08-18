// internal/cache/cache.go
package cache

import (
    "database/sql"
    _ "modernc.org/sqlite" // pure Go driver
    "fmt"
)

type Cache struct {
    db *sql.DB
}

func NewCache(path string) (*Cache, error) {
    db, err := sql.Open("sqlite", path)
    if err != nil { return nil, err }

    // Create table if not exists
    query := `
    CREATE TABLE IF NOT EXISTS inventory (
        product_id TEXT,
        product_name TEXT,
        quantity INTEGER,
        query TEXT,
        last_updated DATETIME
    );`
    _, err = db.Exec(query)
    if err != nil { return nil, err }

    return &Cache{db: db}, nil
}

// Save search results
func (c *Cache) Save(query string, results []string) error {
    tx, _ := c.db.Begin()
    defer tx.Commit()
    for _, item := range results {
        _, err := tx.Exec(`INSERT INTO inventory(product_id, product_name, quantity, query, last_updated)
                           VALUES(?, ?, ?, ?, CURRENT_TIMESTAMP)`,
            item, item, 1, query) // Quantity stub
        if err != nil { return err }
    }
    return nil
}

// Query cache if live search fails
func (c *Cache) Search(query string) ([]string, error) {
    rows, err := c.db.Query(`SELECT product_name FROM inventory WHERE query LIKE ?`, "%"+query+"%")
    if err != nil { return nil, err }
    defer rows.Close()

    var results []string
    for rows.Next() {
        var name string
        _ = rows.Scan(&name)
        results = append(results, name)
    }
    return results, nil
}
