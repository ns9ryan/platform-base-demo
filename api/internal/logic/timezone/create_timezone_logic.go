// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package timezone

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTimezoneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateTimezoneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTimezoneLogic {
	return &CreateTimezoneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateTimezoneLogic) CreateTimezone(req *types.CreateTimezoneRequest) (resp *types.CreateTimezoneResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
