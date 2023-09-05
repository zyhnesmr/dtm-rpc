package wallets

import (
	"errors"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var (
	ErrNotFound     = sqlx.ErrNotFound
	ErrDB           = errors.New("DB error")
	ErrUserNotExist = errors.New("user not exists")
)
