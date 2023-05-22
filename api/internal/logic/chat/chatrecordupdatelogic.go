package chat

import (
	"context"
	"fmt"
	"time"

	"filetool/api/internal/svc"
	"filetool/api/internal/types"
	"filetool/models"
	errpkg "filetool/pkg/error"
	"filetool/pkg/utils"

	log "github.com/pion/ion-log"
	"github.com/zeromicro/go-zero/core/logx"
)

type ChatRecordUpdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewChatRecordUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChatRecordUpdateLogic {
	return &ChatRecordUpdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChatRecordUpdateLogic) ChatRecordUpdate(req *types.ChatRecordUpdateReq) (resp *types.ChatRecordUpdateResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.ChatRecordUpdateResp{Success: "false"}

	if req.AuthKey != "chenlinhua1228" {
		log.Errorf(">>ChatRecordUpdate err, auth failed")
		resp.Errcode = fmt.Sprint(errpkg.NotImplemented)
		resp.ErrMsg = "验证错误"
		return resp, nil
	}

	var info models.ChatRecord
	info.Id = utils.Str2int(req.Id)
	info.Title = req.Title
	info.ChatRecord = req.ChatRecord
	info.UserName = req.UserName
	info.UpdateBy = req.UserName
	info.UpdatedAt = time.Now()

	err = l.svcCtx.TGormCyChatRecordModel.Update(&info)
	if err != nil {
		log.Errorf(">>ChatRecordUpdate err, reason: %v", err)
		resp.Errcode = fmt.Sprint(errpkg.InternalError)
		resp.ErrMsg = "更新记录失败"
		return resp, nil
	}

	resp.Success = "true"
	resp.Errcode = fmt.Sprint(errpkg.Ok)
	return resp, nil
}
