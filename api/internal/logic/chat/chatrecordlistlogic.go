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
	resp = &types.ChatRecordListResp{Success: "false"}
	var fieldStr string
	if req.Title != "" {
		fieldStr = " `title` like '%" + req.Title + "%' "
	}

	list, count, err := l.svcCtx.TGormCyChatRecordModel.List(fieldStr, req.Page, req.Limit)
	if err != nil {
		logx.Errorf("list ChatRecordList error, reason: %s", err)
		resp.Errcode = fmt.Sprint(errpkg.InternalError)
		resp.ErrMsg = "获取记录失败"
		return resp, nil
	}

	//封装返回的记录
	for _, index := range *list {
		var info types.ChatRecordInfo
		info.Id = utils.Int2str(index.Id)
		info.Title = index.Title
		info.ChatRecord = index.ChatRecord
		info.UserName = index.UserName
		info.CreatedAt = utils.FmtTime(index.CreatedAt)

		//处理日期零值
		info.UpdatedAt = utils.FmtTime(index.UpdatedAt)
		if info.UpdatedAt == "0001-01-01 00:00:00" {
			info.UpdatedAt = ""
		}

		info.CreateBy = index.CreateBy
		info.UpdateBy = index.UpdateBy
		resp.Data = append(resp.Data, info)
	}

	//封装响应
	resp.Success = "true"
	resp.Errcode = fmt.Sprint(errpkg.Ok)
	resp.Total = int(count)
	resp.Current = req.Page
	resp.PageSize = len(resp.Data)

	return resp, nil
}
