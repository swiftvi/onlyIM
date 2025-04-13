package logic

import (
	"context"
	"time"

	"kkim/app/user/rpc/internal/svc"
	"kkim/app/user/rpc/user"
	"kkim/pkg/ctx"
	"kkim/pkg/encrypt"
	"kkim/pkg/xcode"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginReq) (*user.LoginResp, error) {
	// todo: add your logic here and delete this line
	if len(in.Phone) < 1 || len(in.Password) < 1 {
		//调用github.com/pkg/errors的WithStack方法，将错误栈信息添加到error中，不会丢失原始错误信息
		return nil, errors.WithStack(xcode.NewError(xcode.REQUEST_PARAM_ERROR, "phone or password is missing"))
	}
	//先找到实体，但这个不是相当于select *？
	userEntity, err := l.svcCtx.UserModels.FindOneByPhone(l.ctx, in.Phone)
	if err != nil {
		return nil, errors.Wrapf(xcode.NewError(xcode.REQUEST_PARAM_ERROR, "failed to find user by phone"), "failed to find user by phone %s: %v", in.Phone, err)
	}
	//是否单独写一个select pwd 比较好？

	if !encrypt.VerifyHashedPwd([]byte(userEntity.Password.String), []byte(in.Password)) {
		return nil, errors.WithStack(xcode.NewError(xcode.REQUEST_PARAM_ERROR, "password is incorrect"))
	}

	token, err := ctx.GenJwtToken(userEntity.Id, l.svcCtx.Config.JWT.Secret, l.svcCtx.Config.JWT.Expire)
	if err != nil {
		return nil, errors.Wrapf(xcode.NewServerCommonError(), "failed to generate jwt token: %v", err)
	}

	exp := l.svcCtx.Config.JWT.Expire + time.Now().Unix()

	return &user.LoginResp{
		Token:  token,
		Expire: exp,
	}, nil
}
