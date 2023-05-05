package chat

import (
	"context"

	"filetool/api/internal/svc"
	"filetool/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRecordListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRecordListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRecordListLogic {
	return &ChatRecordListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRecordListLogic) ChatRecordList(req *types.ChatRecordListReq) (resp *types.ChatRecordListResp, err error) {
	// todo: add your logic here and delete this line

	return
}
