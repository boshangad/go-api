package initialize

import (
	"fmt"
	"github.com/boshangad/go-api/ent"
	"github.com/boshangad/go-api/global"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	// 如果想要不使用nginx代理前端网页，可以修改 web/.env.production 下的
	// VUE_APP_BASE_API = /
	// VUE_APP_BASE_PATH = http://localhost
	// 然后执行打包命令 npm run build。在打开下面4行注释
	//Router.LoadHTMLGlob("./dist/*.html") // npm打包成dist的路径
	//Router.Static("/favicon.ico", "./dist/favicon.ico")
	//Router.Static("/static", "./dist/static")   // dist里面的静态资源
	//Router.StaticFile("/", "./dist/index.html") // 前端网页入口页面

	// 为用户头像和文件提供静态地址
	//Router.StaticFS(global.G_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path))
	// 打开就能玩https了
	// Router.Use(middleware.LoadTls())
	global.G_LOG.Info("use middleware logger")
	// 跨域,如需跨域可以打开
	//Router.Use(middleware.Cors())
	global.G_LOG.Info("use middleware cors")
	//Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.G_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用

	//获取路由组实例
	//systemRouter := router.RouterGroupApp.System
	//exampleRouter := router.RouterGroupApp.Example
	//autocodeRouter := router.RouterGroupApp.Autocode
	//PublicGroup := Router.Group("")
	//{
	//	systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由 不做鉴权
	//	systemRouter.InitInitRouter(PublicGroup) // 自动初始化相关
	//}
	//PrivateGroup := Router.Group("")
	//PrivateGroup.Use(middleware.JWTAuth()).Use(middleware.CasbinHandler())
	//{
	//	systemRouter.InitApiRouter(PrivateGroup)                    // 注册功能api路由
	//	systemRouter.InitJwtRouter(PrivateGroup)                    // jwt相关路由
	//	systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户路由
	//	systemRouter.InitMenuRouter(PrivateGroup)                   // 注册menu路由
	//	systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
	//	systemRouter.InitCasbinRouter(PrivateGroup)                 // 权限相关路由
	//	systemRouter.InitAutoCodeRouter(PrivateGroup)               // 创建自动化代码
	//	systemRouter.InitAuthorityRouter(PrivateGroup)              // 注册角色路由
	//	systemRouter.InitSysDictionaryRouter(PrivateGroup)          // 字典管理
	//	systemRouter.InitSysOperationRecordRouter(PrivateGroup)     // 操作记录
	//	systemRouter.InitSysDictionaryDetailRouter(PrivateGroup)    // 字典详情管理
	//	exampleRouter.InitFileUploadAndDownloadRouter(PrivateGroup) // 文件上传下载功能路由
	//	exampleRouter.InitExcelRouter(PrivateGroup)                 // 表格导入导出
	//	exampleRouter.InitCustomerRouter(PrivateGroup)              // 客户路由
	//
	//	// Code generated by github.com/flipped-aurora/gin-vue-admin/server Begin; DO NOT EDIT.
	//	autocodeRouter.InitSysAutoCodeExampleRouter(PrivateGroup)
	//	// Code generated by github.com/flipped-aurora/gin-vue-admin/server End; DO NOT EDIT.
	//}

	// 安装插件
	//InstallPlugin(PublicGroup, PrivateGroup)

	global.G_LOG.Info("router register success")
	return Router
}


// 无路由执行
func noRoute(c *gin.Context)  {
	c.Abort()
	if c.IsWebsocket() {
		return
	}
	accepts := strings.Split(strings.TrimSpace(c.GetHeader("accept")), ",")
	accept := accepts[0]
	acceptType := strings.Split(accept, "/")[0]
	switch acceptType {
	case "text":
		c.HTML(http.StatusNotFound, "", "")
	case "image":
	case "audio":
	case "video":
	default:
		c.SecureJSON(http.StatusNotFound, map[string]interface{}{"error": http.StatusNotFound, "msg": "Page not found."})
	}
	return
}

// RunServer 运行服务
func RunServer() (e *gin.Engine) {
	gin.SetMode(global.G_CONFIG.System.Env)
	e = gin.New()
	// 未找到路由处理方法
	e.NoRoute(noRoute)
	// 默认的组
	rootGroup := e.Group("/")
	rootGroup.Use(gin.CustomRecovery(func(c *gin.Context, err interface{}) {
		if gin.Mode() == gin.DebugMode {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"error": http.StatusInternalServerError,
				"msg": fmt.Sprintf("%s", err),
			})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": http.StatusInternalServerError,
			"msg": "Service exception, please try again",
		})
		return
	}))
	// 测试环境下定义模型
	if gin.Mode() == gin.DebugMode {
		rootGroup.GET("entviz", gin.WrapH(ent.ServeEntviz()))
	}
	return e
}