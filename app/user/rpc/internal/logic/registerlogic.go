package logic

import (
	"context"
	"database/sql"
	"errors"
	"kkim/app/user/model"
	"kkim/app/user/rpc/internal/svc"
	"kkim/app/user/rpc/user"
	"kkim/pkg/encrypt"
	"kkim/pkg/ctx"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)
var(
	ErrHasBennRegistered = errors.New("The user has been registered")
	ErrInvalidPassword = errors.New("invalid password")
	ErrUnableToRegister = errors.New("unable to register")
	ErrGenerateJwtToken = errors.New("unable to generate jwt token")
)
type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterReq) (*user.RegisterResp, error) {
	// todo: add your logic here and delete this line
	//根据手机号判断用户是否存在
	userEntity, err := l.svcCtx.UserModels.FindOneByPhone(l.ctx, in.Phone)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if userEntity != nil {
		return nil, ErrHasBennRegistered
	}

	//make new user
	userEntity = &model.Users{
		Id: in.Phone,
		Phone: in.Phone,
		Nickname: in.Nickname,
		Avatar: in.Avatar,
		Gender: sql.NullInt64{
			Int64: int64(in.Gender),
			Valid: true,
		},
	}
	if len(in.Password) > 0 {
		hashedPwd, err := encrypt.GenHashedPwd([]byte(in.Password))
		if err != nil {
			return nil, ErrInvalidPassword
		}
		userEntity.Password = sql.NullString{
			String: string(hashedPwd),
			Valid: true,
		}
	} else {
		return nil, ErrInvalidPassword
	}

	//insert user
	_, err = l.svcCtx.UserModels.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, ErrUnableToRegister
	}

	//generate jwt token
	token, err := ctx.GenJwtToken(userEntity.Id, l.svcCtx.Config.JWT.Secret, l.svcCtx.Config.JWT.Expire)
	if err != nil {
		return nil, ErrGenerateJwtToken
	}

	exp := l.svcCtx.Config.JWT.Expire + time.Now().Unix()

	return &user.RegisterResp{
		Token: token,
		Expire: exp,
	}, nil
}
