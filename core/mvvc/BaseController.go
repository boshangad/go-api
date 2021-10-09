package mvvc

import (
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ControllerInterface 接口
type ControllerInterface interface {
}

type Controller struct {
	ControllerInterface
	// 请求对象
	Context *gin.Context
	// 应用用户
	App *ent.App
	// 应用用户
	AppUser *ent.AppUser
	// 用户
	User *ent.User
	// 应用用户登录token
	AppUserToken *ent.AppUserToken
}

// Init 初始化控制器数据
func (that *Controller) Init(c *gin.Context) {
	that.Context = c
	that.App = utils.GetGinApp(that.Context)
	that.AppUserToken = utils.GetGinAppUserToken(that.Context)
	that.AppUser = utils.GetGinAppUser(that.Context)
	that.User = utils.GetGinUser(that.Context)
}

// ShouldBind 关联数据
func (that Controller) ShouldBind(data interface{}) (err error) {
	if that.Context.Request.Method == http.MethodGet {
		err = that.Context.ShouldBindQuery(data)
		if err != nil {
			err = that.Context.ShouldBind(data)
		}
	} else {
		err = that.Context.ShouldBind(data)
	}
	if err != nil {
		return
	}
	// 对字符串进行格式化处理，移除空格
	utils.TrimSpace(data)
	return
}

// JsonOut 输出json数据
func (that Controller) JsonOut(error int64, msg string, data interface{})  {
	response := gin.H{}
	response["error"] = error
	if msg == "" {
		msg = "success"
	}
	response["msg"] = msg
	if data != nil {
		response["data"] = data
	}
	that.Context.AbortWithStatusJSON(http.StatusOK, response)
	return
}

func (that Controller) JsonOutByError(error int64, msg error, data interface{})  {
	that.JsonOut(error, msg.Error(), data)
	return
}