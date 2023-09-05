package wallets

import (
	"database/sql"
	"fmt"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ WalletsModel = (*customWalletsModel)(nil)

type (
	// WalletsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWalletsModel.
	WalletsModel interface {
		walletsModel
		GetRawDB() (*sql.DB, error)
		AddWithTx(tx *sql.Tx, userId int64, money int64) error
		DelWithTx(tx *sql.Tx, userId int64, money int64) error
	}

	customWalletsModel struct {
		*defaultWalletsModel
	}
)

// NewWalletsModel returns a model for the database table.
func NewWalletsModel(conn sqlx.SqlConn) WalletsModel {
	return &customWalletsModel{
		defaultWalletsModel: newWalletsModel(conn),
	}
}

func (m *customWalletsModel) GetRawDB() (*sql.DB, error) {
	return m.conn.RawDB()
}

func (m *customWalletsModel) AddWithTx(tx *sql.Tx, userId int64, money int64) error {
	query := fmt.Sprintf("update %s set wallet = wallet+$1 where user_id = $2", m.table)
	result, err := tx.Exec(query, money, userId)
	if err != nil {
		return ErrDB
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return ErrDB
	}

	if affected == 0 {
		return ErrUserNotExist
	}

	return nil
}

func (m *customWalletsModel) DelWithTx(tx *sql.Tx, userId int64, money int64) error {
	query := fmt.Sprintf("update %s set wallet = wallet-$1 where user_id = $2 and wallet>=$1", m.table)
	result, err := tx.Exec(query, money, userId)
	if err != nil {
		return ErrDB
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return ErrDB
	}

	if affected == 0 {
		return ErrUserNotExist
	}

	return nil
}
