// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package region

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRegionLogic {
	return &CreateRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRegionLogic) CreateRegion(req *types.CreateRegionRequest) (resp *types.CreateRegionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
