// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package language

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ReorderLanguageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewReorderLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ReorderLanguageLogic {
	return &ReorderLanguageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ReorderLanguageLogic) ReorderLanguage(req *types.ReorderLanguageRequest) (resp *types.ReorderLanguageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
