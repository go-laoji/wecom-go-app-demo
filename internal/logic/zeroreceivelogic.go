package logic

import (
	"context"
	"errors"
	"zero-svr/internal/svc"
	"zero-svr/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type ZeroreceiveLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewZeroreceiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) ZeroreceiveLogic {
	return ZeroreceiveLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ZeroreceiveLogic) Zeroreceive(req types.MsgRequest, body []byte) (*types.MsgResponse, error) {
	// todo: add your logic here and delete this line
	echoStr, cryptErr := l.svcCtx.WxBiz.DecryptMsg(req.MsgSignature, req.TimeStamp, req.Nonce, body)
	if cryptErr != nil {
		return &types.MsgResponse{}, errors.New(cryptErr.ErrMsg)
	}
	l.Info(string(echoStr))
	return &types.MsgResponse{ErrCode: 0, ErrMsg: "ok", Data: []byte("SUCCESS")}, nil
}
