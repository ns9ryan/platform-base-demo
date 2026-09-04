// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package region

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllRegionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAllRegionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllRegionsLogic {
	return &ListAllRegionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAllRegionsLogic) ListAllRegions(req *types.ListAllRegionsRequest) (resp *types.ListAllRegionsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
