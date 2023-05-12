package chat

import (
	"context"
	"fmt"

	"filetool/api/internal/svc"
	"filetool/api/internal/types"
	"filetool/models"
	errpkg "filetool/pkg/error"

	log "github.com/pion/ion-log"
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
	resp = &types.ChatRecordAddResp{}
	resp.Success = "false"

	var info models.ChatRecord
	info.Title = req.Title
	info.ChatRecord = req.ChatRecord
	info.UserName = req.UserName
	err = l.svcCtx.TGormCyChatRecordModel.Insert(&info)
	if err != nil {
		log.Errorf(">>ChatRecordAdd err, reason: %v", err)
		resp.Errcode = fmt.Sprint(errpkg.InternalError)
		resp.ErrMsg = "上传记录失败"
		return resp, nil
	}

	resp.Success = "true"
	resp.Errcode = fmt.Sprint(errpkg.Ok)
	return
}
