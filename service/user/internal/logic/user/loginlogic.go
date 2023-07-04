package user

import (
	"context"
	"errors"

	"book/service/user/internal/svc"
	"book/service/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	resp = new(types.LoginResp)
	res, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.Number)
	if err != nil {
		err = errors.New("查询用户出错")
		return
	}
	if res.Password != req.Password {
		err = errors.New("密码错误")
		return
	}
	resp.AccessToken = "123"
	resp.Expiration = "13"
	return
}
