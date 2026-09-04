// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package region

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReorderRegionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReorderRegionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReorderRegionLogic {
	return &ReorderRegionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReorderRegionLogic) ReorderRegion(req *types.ReorderRegionRequest) (resp *types.ReorderRegionResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
