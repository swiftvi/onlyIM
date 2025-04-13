package logic

import (
	"context"
	"errors"

	"kkim/app/user/model"
	"kkim/app/user/rpc/internal/svc"
	"kkim/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *user.FindUserReq) (*user.FindUserResp, error) {
	// todo: add your logic here and delete this line
	if len(in.Ids)==0 && len(in.Phone)==0 && len(in.Name)==0 {
		return nil, errors.New("Invalid params")
	}
	var (
		userEntities []*model.Users
		err error
	)
	
	if in.Phone != "" {
		var userEntity *model.Users
		userEntity, err = l.svcCtx.UserModels.FindOneByPhone(l.ctx, in.Phone)
		logx.Errorf("err: %v", err)
		if err == nil {
			userEntities = append(userEntities, userEntity)
			logx.Infof("userEntity: %v", userEntity)
		}		
	} else if in.Name != "" {
		userEntities, err = l.svcCtx.UserModels.ListByName(l.ctx, in.Name)
		logx.Infof("userEntity: %v", userEntities)
	} else if len(in.Ids)>0 {
		userEntities, err = l.svcCtx.UserModels.ListByIds(l.ctx, in.Ids)
	}
	if err != nil {
		return nil, err
	}
	var resp []*user.UserEntity
	err = copier.Copy(&resp, &userEntities)
	if err != nil {
		return nil, err
	}
	
	//用日志打印出resp结构来检查下
	
	

	return &user.FindUserResp{
		User: resp,
	}, nil
}
