package logic

import (
	"context"

	"czdemo/application/applet/internal/svc"
	"czdemo/application/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AppletLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAppletLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AppletLogic {
	return &AppletLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AppletLogic) Applet(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
