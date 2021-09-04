package middlewares

import (
	"fmt"
	"github.com/boshangad/go-api/core/config"
	"github.com/boshangad/go-api/core/global"
	"github.com/boshangad/go-api/ent"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func init() {
	global.CasbinAuthRequiredLogin.LoadNoAccess(config.Get().NoAccess)
}

// CheckAuth 检查用户权限
func CheckAuth(gh *gin.Context) {
	appUserInterface, ok := gh.Get("AppUser")
	casbinEnforcer := global.CasbinAuthRequiredLogin.Enforcer
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