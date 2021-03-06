package store

import (
	"context"

	"github.com/kingofmidas/gfc-api/internal/model"
)

// Create ...
func (s *Store) Create(order *model.Order) error {
	var orderID int

	ctx := context.Background()
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	err = tx.QueryRow(
		"INSERT INTO orders (status) VALUES ($1) RETURNING id",
		order.Status,
	).Scan(&orderID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, orderItem := range order.ItemList {
		_, err = tx.ExecContext(
			ctx, "INSERT INTO items_orders (item_id, order_id, count) VALUES ($1, $2, $3) RETURNING order_id",
			orderItem.ItemID, orderID, orderItem.Count,
		)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// Update ...
func (s *Store) Update(status string, id int) error {
	err := s.db.QueryRow(
		"UPDATE orders SET status=$1 WHERE id=$2 RETURNING id",
		status,
		id,
	).Scan(&id)
	if err != nil {
		return err
	}

	return nil
}

// Get ...
func (s *Store) Get(status string) (*[]model.Order, error) {
	rows, err := s.db.Query(
		"SELECT id, status FROM orders WHERE status=$1",
		status,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []model.Order

	for rows.Next() {
		var orderID int
		var orderStatus string

		err := rows.Scan(&orderID, &orderStatus)
		if err != nil {
			return nil, err
		}
		orders = append(orders, model.Order{ID: orderID, Status: orderStatus})
	}

	return &orders, nil
}
