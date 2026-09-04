package timezoneservicelogic

import (
	"context"

	"oa.98ent.com/p9/platform-base/rpc/internal/svc"
	"oa.98ent.com/p9/platform-base/rpc/pb/base/timezone"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLogic {
	return &GetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取时区
func (l *GetLogic) Get(in *timezone.GetTimezoneRequest) (*timezone.GetTimezoneResponse, error) {
	// todo: add your logic here and delete this line

	return &timezone.GetTimezoneResponse{}, nil
}
