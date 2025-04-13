package user

import (
	"context"

	"kkim/app/user/api/internal/svc"
	"kkim/app/user/api/internal/types"
	"kkim/app/user/rpc/user"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	// todo: add your logic here and delete this line
	//l.svcCtx.User即为svc中创建的UserRpc客户端，可调用对应服务的方法
	//需要构建user.RegisterReq对象，传递给对应的方法 该返回值是userclient的RegisterResp
	registerResp, err := l.svcCtx.User.Register(l.ctx, &user.RegisterReq{
		Nickname: req.Nickname,
		Password: req.Password,
		Phone:   req.Phone,
		Avatar: req.Avatar,
		Gender: int32(req.Gender),
	})
	if err != nil {
		return nil, err
	}

	var res types.RegisterResp
	_ = copier.Copy(&res, &registerResp)
	return &res, nil
}
