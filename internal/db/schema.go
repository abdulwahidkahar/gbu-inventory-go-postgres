package db

import "database/sql"

func EnsureSchema(conn *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS inventory_items (
		id BIGSERIAL PRIMARY KEY,
		sku VARCHAR(50) NOT NULL UNIQUE,
		name VARCHAR(255) NOT NULL,
		category VARCHAR(100) NOT NULL,
		supplier_name VARCHAR(255) NOT NULL,
		purchase_price NUMERIC(14,2) NOT NULL CHECK (purchase_price >= 0),
		selling_price NUMERIC(14,2) NOT NULL CHECK (selling_price >= 0),
		stock INTEGER NOT NULL CHECK (stock >= 0),
		reorder_level INTEGER NOT NULL CHECK (reorder_level >= 0),
		is_active BOOLEAN NOT NULL DEFAULT TRUE,
		created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
		updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
	);
	`

	_, err := conn.Exec(query)
	return err
}
