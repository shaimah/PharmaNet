package cache

import (
	"database/sql"
	_ "modernc.org/sqlite"
)

type Store struct {
	DB *sql.DB
}

func NewStore(path string) (*Store, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, err
	}
	// Create table if not exists
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS stock (
		product_id TEXT,
		product_name TEXT,
		strength TEXT,
		pack_size TEXT,
		location TEXT,
		quantity INTEGER,
		expiry_date TEXT,
		last_updated TEXT
	);`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		return nil, err
	}
	return &Store{DB: db}, nil
}

func (s *Store) Save(records []connectors.StockRecord) error {
	tx, _ := s.DB.Begin()
	stmt, _ := tx.Prepare(`
	INSERT INTO stock(product_id, product_name, strength, pack_size, location, quantity, expiry_date, last_updated)
	VALUES(?, ?, ?, ?, ?, ?, ?, ?)
	ON CONFLICT(product_id) DO UPDATE SET
	product_name=excluded.product_name,
	quantity=excluded.quantity,
	last_updated=excluded.last_updated
	`)
	defer stmt.Close()
	for _, r := range records {
		_, _ = stmt.Exec(r.ProductID, r.ProductName, r.Strength, r.PackSize, r.Location, r.Quantity, r.ExpiryDate, r.LastUpdated)
	}
	return tx.Commit()
}

func (s *Store) Query(query string, limit int) ([]connectors.StockRecord, error) {
	rows, err := s.DB.Query("SELECT * FROM stock WHERE LOWER(product_name) LIKE ? LIMIT ?", "%"+query+"%", limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var out []connectors.StockRecord
	for rows.Next() {
		var r connectors.StockRecord
		_ = rows.Scan(&r.ProductID, &r.ProductName, &r.Strength, &r.PackSize, &r.Location, &r.Quantity, &r.ExpiryDate, &r.LastUpdated)
		out = append(out, r)
	}
	return out, nil
}
