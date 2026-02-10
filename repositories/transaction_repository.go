package repositories

import (
	"cashier-api/models"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

type TransactionRepository struct {
	db *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{db: db}
}

func (repo *TransactionRepository) CreateTransaction(items []models.CheckoutItem) (*models.Transaction, error) {
	tx, err := repo.db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	totalAmount := 0
	details := make([]models.TransactionDetail, 0)

	for _, item := range items {
		var productPrice, stock int
		var productName string

		err := tx.QueryRow("SELECT name, price, stock FROM products WHERE id = $1", item.ProductID).Scan(&productName, &productPrice, &stock)
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("product id %d not found", item.ProductID)
		}
		if err != nil {
			return nil, err
		}

		subtotal := productPrice * item.Quantity
		totalAmount += subtotal

		_, err = tx.Exec("UPDATE products SET stock = stock - $1 WHERE id = $2", item.Quantity, item.ProductID)
		if err != nil {
			return nil, err
		}

		details = append(details, models.TransactionDetail{
			ProductID:   item.ProductID,
			ProductName: productName,
			Quantity:    item.Quantity,
			Subtotal:    subtotal,
		})
	}

	var transactionID int
	var createdAt time.Time
	err = tx.QueryRow("INSERT INTO transactions (total_amount) VALUES ($1) RETURNING id, created_at", totalAmount).Scan(&transactionID, &createdAt)
	if err != nil {
		return nil, err
	}

	for i := range details {
		var transactionDetailID int

		details[i].TransactionID = transactionID

		err = tx.QueryRow("INSERT INTO transaction_details (transaction_id, product_id, quantity, subtotal) VALUES ($1, $2, $3, $4) RETURNING id", transactionID, details[i].ProductID, details[i].Quantity, details[i].Subtotal).Scan(&transactionDetailID)
		if err != nil {
			return nil, err
		}

		details[i].ID = transactionDetailID
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return &models.Transaction{
		ID:          transactionID,
		CreatedAt:   createdAt,
		TotalAmount: totalAmount,
		Details:     details,
	}, nil
}

func (repo *TransactionRepository) GetReport() (*models.Report, error) {
	var report models.Report

	query := "SELECT COALESCE(SUM(total_amount), 0) AS total_amount, COUNT(*) AS total_transaction FROM transactions WHERE DATE(created_at) = CURRENT_DATE"
	err := repo.db.QueryRow(query).Scan(&report.TotalRevenue, &report.TotalTransaction)
	if err == sql.ErrNoRows {
		return nil, errors.New("Report not found")
	}

	queryTopProduct := `
		SELECT
			p.name,
			SUM(td.quantity) AS total_qty
		FROM transaction_details td
		JOIN transactions t ON t.id = td.transaction_id
		JOIN products p ON td.product_id = p.id
		WHERE DATE(t.created_at) = CURRENT_DATE
		GROUP BY td.product_id, p.name
		ORDER BY total_qty DESC
		LIMIT 1
	`
	err = repo.db.QueryRow(queryTopProduct).Scan(&report.TopSellProduct.Name, &report.TopSellProduct.QuantitySell)
	if err == sql.ErrNoRows {
		return nil, errors.New("Report not found")
	}

	if err != nil {
		return nil, err
	}

	return &report, nil
}
