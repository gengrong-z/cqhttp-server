package handler

import (
	"cqhttp-server/config"
	"cqhttp-server/internal/pkg/wscore"
)

func Rank(ctx *wscore.Context) error {
	SetPicture(ctx, config.SavePath, config.Static.PixivUrl)
	return nil
}

func Sketchy(ctx *wscore.Context) error {
	SetPicture(ctx, config.WHPath, config.Static.WHUrl)
	return nil
}
