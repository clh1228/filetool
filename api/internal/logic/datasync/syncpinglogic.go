package logic

import (
	"context"

	"filetool/api/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SyncPingLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSyncPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) SyncPingLogic {
	return SyncPingLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SyncPingLogic) SyncPing() error {
	// todo: add your logic here and delete this line

	return nil
}
