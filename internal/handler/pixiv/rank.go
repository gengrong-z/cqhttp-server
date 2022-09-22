package pixiv

import (
	"cqhttp-server/config"
	"cqhttp-server/pkg"
	"cqhttp-server/pkg/wscore"
	"fmt"
)

func Rank(ctx *wscore.Context) error {
	path := pkg.GetRandFileAbsPath(config.SavePath)
	if path == "" {
		pkg.PixivCraw(config.Static.PixivUrl)
		path = pkg.GetRandFileAbsPath(config.SavePath)
	}
	callback := &wscore.Callback{
		Params: &wscore.CallbackSender{
			Message: fmt.Sprintf("[CQ:image,file=%s]", path),
		},
	}
	ctx.Callback = callback
	return nil
}
