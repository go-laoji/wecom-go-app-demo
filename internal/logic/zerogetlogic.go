package logic

import (
	"context"
	"errors"

	"zero-svr/internal/svc"
	"zero-svr/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ZerogetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZerogetLogic(ctx context.Context, svcCtx *svc.ServiceContext) ZerogetLogic {
	return ZerogetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZerogetLogic) Zeroget(req types.MsgRequest) (*types.MsgResponse, error) {
	// todo: add your logic here and delete this line
	echoStr, cryptErr := l.svcCtx.WxBiz.VerifyURL(req.MsgSignature, req.TimeStamp, req.Nonce, req.EchoStr)
	if cryptErr != nil {
		return &types.MsgResponse{}, errors.New(cryptErr.ErrMsg)
	}
	return &types.MsgResponse{ErrCode: 0, ErrMsg: "ok", Data: echoStr}, nil
}
