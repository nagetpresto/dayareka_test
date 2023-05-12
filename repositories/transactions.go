package repositories

import (
	"BE/models"
	"database/sql"
)

type TransactionRepository interface {
	GetOneTransaction(ID int) (models.Transaction, error)
	PostTransaction(transaction models.Transaction) (models.Transaction, error)
	GetTransactionsSortedByLatest(menu string, price int) ([]models.Transaction, error)
}

func RepositoryTransaction(db *sql.DB) *repository {
	return &repository{db}
}

func (r *repository) GetOneTransaction(ID int) (models.Transaction, error) {
	var transaction models.Transaction
	err := r.db.QueryRow(`
		SELECT *
		FROM transactions
		WHERE id = $1`,
		ID).Scan(
		&transaction.ID,
		&transaction.UserID,
		&transaction.Menu,
		&transaction.Price,
		&transaction.Qty,
		&transaction.Total,
		&transaction.Payment,
		&transaction.CreatedAt,
	)
	if err != nil {
		return transaction, err
	}

	// Retrieve user details
	var user models.User
	err = r.db.QueryRow(`
		SELECT *
		FROM users
		WHERE id = $1`,
		transaction.UserID,
	).Scan(
		&user.ID,
		&user.Name,
	)
	if err != nil {
		return transaction, err
	}
	transaction.User = user
	return transaction, err
}

func (r *repository) PostTransaction(transaction models.Transaction) (models.Transaction, error) {
	_, err := r.db.Exec(`
		INSERT INTO transactions (user_id, menu, price, qty, payment, total, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)`,
		transaction.UserID, transaction.Menu, transaction.Price, transaction.Qty,
		transaction.Payment, transaction.Total, transaction.CreatedAt,
	)
	if err != nil {
		return transaction, err
	}

	var lastInsertedID int
	err = r.db.QueryRow("SELECT lastval()").Scan(&lastInsertedID)
	if err != nil {
		return transaction, err
	}

	transaction.ID = int(lastInsertedID)
	
	return transaction, err
}

func (r *repository) GetTransactionsSortedByLatest(menu string, price int) ([]models.Transaction, error) {
	var transactions []models.Transaction
	var params []interface{}
	queryStr := `
		SELECT t.*, u.name AS customer_name
		FROM transactions t
		LEFT JOIN users u ON t.user_id = u.id
		WHERE 1 = 1`

	if menu != "" {
		queryStr += ` AND t.menu LIKE '%' || $1 || '%'`
		params = append(params, menu)
	} else if price != 0 {
		queryStr += ` AND t.price = $1`
		params = append(params, price)
	}

	queryStr += ` ORDER BY`
	if menu != "" || price != 0 {
		queryStr += ` u.name ASC, t.created_at DESC`
	} else {
		queryStr += ` t.created_at DESC`
	}

	rows, err := r.db.Query(queryStr, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(
			&transaction.ID,
			&transaction.UserID,
			&transaction.Menu,
			&transaction.Price,
			&transaction.Qty,
			&transaction.Total,
			&transaction.Payment,
			&transaction.CreatedAt,
			&transaction.User.Name,
		)
		if err != nil {
			return nil, err
		}

		// Retrieve user details
		var user models.User
		err = r.db.QueryRow(`
			SELECT *
			FROM users
			WHERE id = $1`,
			transaction.UserID,
		).Scan(
			&user.ID,
			&user.Name,
		)
		if err != nil {
			return nil, err
		}

		transaction.User = user
		transactions = append(transactions, transaction)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return transactions, nil
}
