package user

import (
	"context"
	"kkim/app/user/api/internal/svc"
	"kkim/app/user/api/internal/types"
	"kkim/app/user/rpc/user"
	"kkim/pkg/ctx"

	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	// todo: add your logic here and delete this line
	uid := ctx.GetUidFromToken(l.ctx)
	infoResp, err := l.svcCtx.User.GetUserInfo(l.ctx, &user.GetUserInfoReq{
		Id: uid,
	})
	if err != nil {
		return nil, err
	}
	var res types.User
	_ = copier.Copy(&res, &infoResp.User)
	return &types.UserInfoResp{
		Info: res,
	}, nil
}
