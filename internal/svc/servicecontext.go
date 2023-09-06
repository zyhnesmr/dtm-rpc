package svc

import (
	"dtm-rpc/internal/config"
	"dtm-rpc/model/wallets"

	"github.com/dtm-labs/dtm/client/dtmcli/dtmimp"
	"github.com/zeromicro/go-zero/core/stores/postgres"
)

type ServiceContext struct {
	Config      config.Config
	WalletModel wallets.WalletsModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	dtmimp.SetCurrentDBType(dtmimp.DBTypePostgres)
	return &ServiceContext{
		Config:      c,
		WalletModel: wallets.NewWalletsModel(postgres.New(c.Datasource)),
	}
}
