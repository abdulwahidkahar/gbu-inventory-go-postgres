package inventory

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) List() ([]Item, error) {
	rows, err := r.db.Query(`
		SELECT id, sku, name, category, supplier_name, purchase_price, selling_price,
		       stock, reorder_level, is_active, created_at, updated_at
		FROM inventory_items
		ORDER BY created_at DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	items := []Item{}
	for rows.Next() {
		var item Item
		if err := rows.Scan(
			&item.ID,
			&item.SKU,
			&item.Name,
			&item.Category,
			&item.SupplierName,
			&item.PurchasePrice,
			&item.SellingPrice,
			&item.Stock,
			&item.ReorderLevel,
			&item.IsActive,
			&item.CreatedAt,
			&item.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, item)
	}

	return items, rows.Err()
}

func (r *Repository) Find(id int64) (*Item, error) {
	var item Item

	err := r.db.QueryRow(`
		SELECT id, sku, name, category, supplier_name, purchase_price, selling_price,
		       stock, reorder_level, is_active, created_at, updated_at
		FROM inventory_items
		WHERE id = $1
	`, id).Scan(
		&item.ID,
		&item.SKU,
		&item.Name,
		&item.Category,
		&item.SupplierName,
		&item.PurchasePrice,
		&item.SellingPrice,
		&item.Stock,
		&item.ReorderLevel,
		&item.IsActive,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *Repository) Create(payload Payload) (*Item, error) {
	var item Item

	err := r.db.QueryRow(`
		INSERT INTO inventory_items
			(sku, name, category, supplier_name, purchase_price, selling_price, stock, reorder_level, is_active)
		VALUES
			($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, sku, name, category, supplier_name, purchase_price, selling_price,
		          stock, reorder_level, is_active, created_at, updated_at
	`,
		payload.SKU,
		payload.Name,
		payload.Category,
		payload.SupplierName,
		payload.PurchasePrice,
		payload.SellingPrice,
		payload.Stock,
		payload.ReorderLevel,
		payload.IsActive,
	).Scan(
		&item.ID,
		&item.SKU,
		&item.Name,
		&item.Category,
		&item.SupplierName,
		&item.PurchasePrice,
		&item.SellingPrice,
		&item.Stock,
		&item.ReorderLevel,
		&item.IsActive,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *Repository) Update(id int64, payload Payload) (*Item, error) {
	var item Item

	err := r.db.QueryRow(`
		UPDATE inventory_items
		SET sku = $2,
		    name = $3,
		    category = $4,
		    supplier_name = $5,
		    purchase_price = $6,
		    selling_price = $7,
		    stock = $8,
		    reorder_level = $9,
		    is_active = $10,
		    updated_at = NOW()
		WHERE id = $1
		RETURNING id, sku, name, category, supplier_name, purchase_price, selling_price,
		          stock, reorder_level, is_active, created_at, updated_at
	`,
		id,
		payload.SKU,
		payload.Name,
		payload.Category,
		payload.SupplierName,
		payload.PurchasePrice,
		payload.SellingPrice,
		payload.Stock,
		payload.ReorderLevel,
		payload.IsActive,
	).Scan(
		&item.ID,
		&item.SKU,
		&item.Name,
		&item.Category,
		&item.SupplierName,
		&item.PurchasePrice,
		&item.SellingPrice,
		&item.Stock,
		&item.ReorderLevel,
		&item.IsActive,
		&item.CreatedAt,
		&item.UpdatedAt,
	)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *Repository) Delete(id int64) (bool, error) {
	result, err := r.db.Exec(`DELETE FROM inventory_items WHERE id = $1`, id)
	if err != nil {
		return false, err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return false, err
	}

	return affected > 0, nil
}
