package user

import (
	"context"

	"kkim/app/user/api/internal/svc"
	"kkim/app/user/api/internal/types"
	"kkim/app/user/rpc/user"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	// todo: add your logic here and delete this line
	loginResp, err := l.svcCtx.User.Login(l.ctx, &user.LoginReq{
		Phone: req.Phone,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	var res types.LoginResp
	_ = copier.Copy(&res, &loginResp)
	return &res, nil
}
