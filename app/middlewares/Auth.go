package middlewares

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/gin-gonic/gin"
	"github.com/tend/wechatServer/core/config"
	"github.com/tend/wechatServer/core/global"
	"github.com/tend/wechatServer/ent"
	"log"
	"net/http"
)

// casbin鉴权用户是否需要进行登录
var casbinEnforcer *casbin.Enforcer

func init() {
	// 检查是否需要用户登录
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && regexMatch(r.act, p.act)
`)
	if err != nil {
		log.Fatalf("error: model: %s", err)
	}
	casbinEnforcer, err = casbin.NewEnforcer(m)
	if err != nil {
		log.Fatalf("error: casbin call newEnforcer: %s", err)
	}

	noAccess := config.Get().NoAccess
	if noAccess != nil && len(noAccess) > 0 {
		var accessItems [][]string
		for access, methods := range noAccess {
			if method, ok := methods.(string); ok {
				accessItems = append(accessItems, []string{"guest", "/"+access, method})
			} else if methodItems, ok := methods.([]string); ok {
				for _, method := range methodItems {
					accessItems = append(accessItems, []string{"guest", "/"+access, method})
				}
			}
		}
		_, _ = casbinEnforcer.AddPolicies(accessItems)
	}
}

// CheckAuth 检查用户权限
func CheckAuth(gh *gin.Context)  {
	appUserInterface, ok := gh.Get("AppUser")
	if !ok {
		enforce, err := casbinEnforcer.Enforce("guest", gh.Request.URL.Path, gh.Request.Method)
		if err != nil {
			log.Println("casbin 检查不通过请稍后再试", err)
		}
		if !enforce {
			gh.AbortWithStatusJSON(http.StatusForbidden, global.JsonResponse{
				Error: http.StatusForbidden,
				Msg: "Authentication failed, API needs to verify your account.",
			})
		}
		return
	}
	// 检查是否存在appUser实体
	appUserModel, ok := appUserInterface.(*ent.AppUser)
	if !ok {
		gh.AbortWithStatusJSON(http.StatusInternalServerError, global.JsonResponse{
			Error: global.ErrMissLoginParams,
			Msg: "Service abnormal, please try again later.",
		})
		return
	}
	// 检查用户是否拥有权限访问
	enforce, err := casbinEnforcer.Enforce(
		fmt.Sprintf("user%d", appUserModel.ID),
		gh.Request.URL.Path,
		gh.Request.Method,
	)
	if err != nil {
		log.Println("casbin 检查应用用户权限异常", err)
	}
	if !enforce {
		//gh.AbortWithStatusJSON(http.StatusOK, global.JsonResponse{
		//	Error: global.ErrRequiredLogin,
		//	Msg: "No permission to access this resource.",
		//})
		//return
	}
}
