package chat

import (
	"context"
	"fmt"

	"filetool/api/internal/svc"
	"filetool/api/internal/types"
	errpkg "filetool/pkg/error"
	"filetool/pkg/utils"

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
	resp = &types.ChatRecordDelResp{}
	resp.Success = "false"

	err = l.svcCtx.TGormCyChatRecordModel.Delete(utils.Str2int(req.Id))
	if err != nil {
		resp.Errcode = fmt.Sprint(errpkg.InternalError)
		resp.ErrMsg = "删除记录失败"
		return resp, nil
	}

	resp.Success = "true"
	resp.Errcode = fmt.Sprint(errpkg.Ok)
	return resp, nil
}
