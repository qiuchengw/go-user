package app

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"github.com/qiuchengw/go-user/config"
	"github.com/qiuchengw/go-user/frontend/errors"
	"github.com/qiuchengw/go-user/frontend/session"
	"github.com/qiuchengw/go-user/frontend/token"
	"github.com/qiuchengw/go-user/util"
)

// 获取请求用户授权的参数(appid, state, scope)
func AuthParaHandler(ctx *gin.Context) {
	// MustAuthHandler(ctx)
	tk := ctx.MustGet("sso_token").(*token.Token)
	ss := ctx.MustGet("sso_session").(*session.Session)

	ss.OAuth2State = string(util.NewRandom())
	if err := session.Set(tk.SessionId, ss); err != nil {
		glog.Errorln(err)
		ctx.JSON(200, errors.ErrInternalServerError)
		return
	}

	resp := struct {
		*errors.Error
		AppId string `json:"appid"`
		State string `json:"state"`
		Scope string `json:"scope"`
	}{
		Error: errors.ErrOK,
		AppId: config.ConfigData.Weixin.Open.App.AppId,
		State: ss.OAuth2State,
		Scope: "snsapi_userinfo",
	}
	ctx.JSON(200, &resp)
	return
}
