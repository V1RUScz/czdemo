package logic

import (
	"context"
	"errors"
	"strings"

	"czdemo/application/applet/internal/code"
	"czdemo/application/applet/internal/svc"
	"czdemo/application/applet/internal/types"
	"czdemo/application/user/rpc/user"
	"czdemo/pkg/encrypt"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	req.Name = strings.TrimSpace(req.Name)
	req.Mobile = strings.TrimSpace(req.Mobile)
	if len(req.Mobile) == 0 {
		return nil, code.RegisterMobileEmpty
	}
	req.Password = strings.TrimSpace(req.Password)
	if len(req.Password) == 0 {
		return nil, code.RegisterPasswordEmpty
	} else {
		req.Password = encrypt.EncPassword(req.Password)
	}
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)
	if len(req.VerificationCode) == 0 {
		return nil, code.VerificationCodeEmpty
	}
	e := checkVerficationCode(req.Mobile,req.VerificationCode,l.svcCtx.BizRedis)
	if e != nil {
		logx.Errorf("EncMobile mobile: %s error: %v", req.Mobile, e)
		return nil, e
	}
	mobile, err := encrypt.EncMobile(req.Mobile)
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &user.FindByMobileRequest{
		Mobile: mobile,
	})
	if err != nil {
		logx.Errorf("FindByMobile error: %v", err)
		return nil, err
	}
	if u != nil && u.UserId > 0 {
		return nil, code.MobileHasRegistered
	}
}

func checkVerficationCode(mobile, verfication_code string, res *redis.Redis) error {
	cachecode,err := getActivationCache(mobile,res)
	if err!= nil{
		return err
	}
	if(cachecode == ""){
		return errors.New("verification code expired")
	}
	if(cachecode != verfication_code){
		return errors.New("verification code failed")
	}
	return nil
}
