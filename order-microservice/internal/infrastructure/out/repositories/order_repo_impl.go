package repositories

import (
	"context"
	"database/sql"

	"github.com/trng-tr/order-microservice/internal/infrastructure/out/models"
)

// OrderRepoImpl implements OrderRepo
type OrderRepoImpl struct {
	db *sql.DB
}

// injection par construteur
func NewOrderRepoImpl(db *sql.DB) *OrderRepoImpl {
	return &OrderRepoImpl{db: db}
}

// Save implement OrderRepo
func (o *OrderRepoImpl) Save(ctx context.Context, model models.OrderModel) (models.OrderModel, error) {
	query := `INSERT INTO orders (customer_id,created_at,status)
	VALUES ($1,$2,$3)
	RETURNING id`
	if err := o.db.QueryRowContext(
		ctx,
		query,
		model.CustomerID,
		model.CreatedAt,
		model.Status,
	).Scan(&model.ID); err != nil {
		return models.OrderModel{}, err
	}
	return model, nil
}

// FindByID implement OrderRepo
func (o *OrderRepoImpl) FindByID(ctx context.Context, id int64) (models.OrderModel, error) {
	query := `SELECT id,customer_id,created_at,status FROM orders WHERE id=$1`
	var model models.OrderModel
	if err := o.db.QueryRowContext(ctx, query, id).Scan(
		&model.ID, &model.CustomerID, &model.CreatedAt, &model.Status,
	); err != nil {
		return models.OrderModel{}, err
	}

	return model, nil
}

// FindAll implement OrderRepo
func (o *OrderRepoImpl) FindAll(ctx context.Context) ([]models.OrderModel, error) {
	query := "SELECT id,customer_id,created_at,status FROM orders"
	rows, err := o.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	var data []models.OrderModel = make([]models.OrderModel, 0)
	for rows.Next() {
		var model models.OrderModel
		if err := rows.Scan(&model.ID, &model.CustomerID, &model.CreatedAt, &model.Status); err != nil {
			return nil, err
		}
		data = append(data, model)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}

func (o *OrderRepoImpl) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM orders WHERE id=$1"
	results, err := o.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := results.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return sql.ErrNoRows
	}
	return nil
}
