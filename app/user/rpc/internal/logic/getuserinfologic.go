package logic

import (
	"context"
	"errors"
	"github.com/jinzhu/copier"
	"kkim/app/user/model"
	"kkim/app/user/rpc/internal/svc"
	"kkim/app/user/rpc/user"
	"github.com/zeromicro/go-zero/core/logx"
)
var (
	ErrInvalidParams   = errors.New("invalid params")
	ErrUserNotFound    = errors.New("User not found")
	ErrUserPwdNotMatch = errors.New("User Pwd not macth")
)
type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	// todo: add your logic here and delete this line
	if len(in.Id)<1 {
		return nil, ErrInvalidParams
	}
	userEntity, err := l.svcCtx.UserModels.FindOne(l.ctx, in.Id)
	if err != nil || err == model.ErrNotFound {
		return nil, ErrUserNotFound
	}

	//copy userEntity to GetUserInfoResp
	//使用第三方库copier
	var resp user.UserEntity

	err = copier.Copy(&resp, userEntity)
	if err != nil {
		return nil, err
	}
	return &user.GetUserInfoResp{
		User: &resp,
	}, nil
}
