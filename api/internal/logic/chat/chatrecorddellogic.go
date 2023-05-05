package chat

import (
	"context"

	"filetool/api/internal/svc"
	"filetool/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRecordDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRecordDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRecordDelLogic {
	return &ChatRecordDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRecordDelLogic) ChatRecordDel(req *types.ChatRecordDelReq) (resp *types.ChatRecordDelResp, err error) {
	// todo: add your logic here and delete this line

	return
}
