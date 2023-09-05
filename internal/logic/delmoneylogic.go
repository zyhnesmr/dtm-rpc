package logic

import (
	"context"
	"database/sql"
	"dtm-rpc/busis"
	"dtm-rpc/internal/svc"
	"dtm-rpc/model/wallets"

	"github.com/dtm-labs/dtm/client/dtmcli"
	"github.com/dtm-labs/dtm/client/dtmgrpc"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type DelMoneyLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDelMoneyLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DelMoneyLogic {
	return &DelMoneyLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DelMoneyLogic) DelMoney(in *busis.AddReq) (*busis.Empty, error) {
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
		return l.svcCtx.WalletModel.DelWithTx(tx, in.UserId, in.Money)
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
