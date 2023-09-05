package logic

import (
	"context"
	"database/sql"
	"dtm-rpc/model/wallets"

	"dtm-rpc/busis"
	"dtm-rpc/internal/svc"

	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AddMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddMoneyLogic {
	return &AddMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddMoneyLogic) AddMoney(in *busis.AddReq) (*busis.Empty, error) {
	//barrier := utils.MustBarrierFromGrpc(l.ctx)
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(codes.Code(1), "xx")
	}

	db, err := l.svcCtx.WalletModel.GetRawDB()
	if err != nil {
		//不返回dtm的错误码 让dtm一直重试
		return nil, status.Error(codes.Code(1), "xx")
	}

	if err = barrier.CallWithDB(db, func(tx *sql.Tx) error {
		return l.svcCtx.WalletModel.AddWithTx(tx, in.UserId, in.Money)
	}); err != nil {
		switch err {
		case nil:
			return &busis.Empty{}, nil
		case wallets.ErrDB:
			return nil, status.Error(codes.Code(1), "xx")
		case wallets.ErrUserNotExist:
			return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
		}
	}

	return &busis.Empty{}, nil
}
