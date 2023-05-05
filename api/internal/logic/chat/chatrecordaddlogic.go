package chat

import (
	"context"

	"filetool/api/internal/svc"
	"filetool/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRecordAddLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRecordAddLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRecordAddLogic {
	return &ChatRecordAddLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRecordAddLogic) ChatRecordAdd(req *types.ChatRecordAddReq) (resp *types.ChatRecordAddResp, err error) {
	// todo: add your logic here and delete this line

	return
}
