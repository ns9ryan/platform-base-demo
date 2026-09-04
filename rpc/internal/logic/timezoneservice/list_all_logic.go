package timezoneservicelogic

import (
	"context"

	"oa.98ent.com/p9/platform-base/rpc/internal/svc"
	"oa.98ent.com/p9/platform-base/rpc/pb/base/timezone"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListAllLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllLogic {
	return &ListAllLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取全部时区
func (l *ListAllLogic) ListAll(in *timezone.ListAllTimezonesRequest) (*timezone.ListAllTimezonesResponse, error) {
	// todo: add your logic here and delete this line

	return &timezone.ListAllTimezonesResponse{}, nil
}
