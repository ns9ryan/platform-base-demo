// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.2

package language

import (
	"context"

	"oa.98ent.com/p9/platform-base/api/internal/svc"
	"oa.98ent.com/p9/platform-base/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListAllLanguagesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListAllLanguagesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListAllLanguagesLogic {
	return &ListAllLanguagesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListAllLanguagesLogic) ListAllLanguages(req *types.ListAllLanguagesRequest) (resp *types.ListAllLanguagesResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
