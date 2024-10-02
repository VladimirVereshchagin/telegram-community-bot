package common

import (
	"database/sql"
)

// UnitOfWork интерфейс для управления транзакциями
type UnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	GetTx() *sql.Tx
}

// unitOfWork реализация интерфейса UnitOfWork
type unitOfWork struct {
	db *sql.DB
	tx *sql.Tx
}

// NewUnitOfWork создает новый экземпляр unitOfWork
func NewUnitOfWork(db *sql.DB) UnitOfWork {
	return &unitOfWork{
		db: db,
	}
}

// Begin начинает новую транзакцию
func (u *unitOfWork) Begin() error {
	tx, err := u.db.Begin()
	if err != nil {
		return err
	}
	u.tx = tx
	return nil
}

// Commit подтверждает транзакцию
func (u *unitOfWork) Commit() error {
	if u.tx == nil {
		return sql.ErrTxDone
	}
	return u.tx.Commit()
}

// Rollback откатывает транзакцию
func (u *unitOfWork) Rollback() error {
	if u.tx == nil {
		return sql.ErrTxDone
	}
	return u.tx.Rollback()
}

// GetTx возвращает текущую транзакцию
func (u *unitOfWork) GetTx() *sql.Tx {
	return u.tx
}
