// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package language

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLanguageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetLanguageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLanguageLogic {
	return &GetLanguageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetLanguageLogic) GetLanguage(req *types.GetLanguageRequest) (resp *types.GetLanguageResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
