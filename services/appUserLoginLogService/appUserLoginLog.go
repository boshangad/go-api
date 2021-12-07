package appUserLoginLogService

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/boshangad/v1/app/global"
	"github.com/boshangad/v1/ent"
	"github.com/boshangad/v1/ent/appuserloginlog"
	"go.uber.org/zap"
)

// 检测和创建一条用户登录日志
// @param uint64 $appId 登录的应用ID
// @param string $ip 登录IP
// @param *http.Request $httpRequest 请求
func CheckAndCreateLoginLogByConfirm(appId uint64, ip string, httpRequest *http.Request) (*ent.AppUserLoginLog, error) {
	var (
		ctx                = context.Background()
		httpDumpRequest, _ = httputil.DumpRequest(httpRequest, false)
	)
	appUserLoginLog, err := global.Db.AppUserLoginLog.Create().
		SetAppID(appId).
		SetIP(ip).
		SetUserAgent(httpRequest.UserAgent()).
		SetClientVersion(httpRequest.UserAgent()).
		SetLoginTypeID(appuserloginlog.LoginTypeUnknow).
		SetContent(string(httpDumpRequest)).
		SetStatus(appuserloginlog.StatusWaitConfirm).
		Save(ctx)
	if err != nil {
		if !ent.IsNotFound(err) {
			global.Log.Warn("create appUserLoginLog failed", zap.Error(err))
		}
		return nil, fmt.Errorf("login failed, please try again later")
	}
	return appUserLoginLog, nil
}
