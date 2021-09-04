package global

import (
	"github.com/boshangad/go-api/utils"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"log"
	"net/http"
	"sort"
	"strings"
)

type casbinEnforcerWithLogin struct {
	Enforcer *casbin.Enforcer
}

var (
	// CasbinAuthRequiredLogin 检查用户是否需要登录
	CasbinAuthRequiredLogin casbinEnforcerWithLogin
)

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
	e, err := casbin.NewEnforcer(m)
	if err != nil {
		log.Fatalf("error: casbin call newEnforcer: %s", err)
	}
	CasbinAuthRequiredLogin = casbinEnforcerWithLogin{
		Enforcer: e,
	}
}

// LoadNoAccess 加载检查器
func (that *casbinEnforcerWithLogin) LoadNoAccess(noAccessItems map[string]interface{})  {
	var (
		err error
		validMethods = []string{
			http.MethodGet, http.MethodHead, http.MethodPost, http.MethodPut,
			http.MethodPatch, http.MethodDelete, http.MethodConnect, http.MethodOptions,
			http.MethodTrace,
		}
		accessItems [][]string
		appendMethod func(access, method string)
		casbinEnforcer = that.Enforcer
	)
	noAccess := noAccessItems
	if noAccess == nil || len(noAccess) < 1 {
		noAccess = map[string]interface{} {}
	}
	appendMethod = func (access, method string) {
		method = strings.ToUpper(strings.TrimSpace(method))
		if method == "ANY" || method == "*" {
			method = ""
		}
		if access[:1] != "/" {
			access = "/" + access
		}
		if method == "" {
			for _, m := range validMethods {
				appendMethod(access, m)
			}
			return
		}
		if !utils.InArrayWithString(method, validMethods) {
			return
		}
		accessItems = append(accessItems, []string{
			"guest",
			access,
			method,
		})
	}
	// 循环加载
	for access, methods := range noAccess {
		if method, ok := methods.(string); ok {
			appendMethod(access, method)
		} else if methodItems, ok := methods.([]string); ok {
			sort.Strings(methodItems)
			for _, method := range methodItems {
				appendMethod(access, method)
			}
		}
	}
	casbinEnforcer.ClearPolicy()
	_, err = casbinEnforcer.AddPolicies(accessItems)
	if err != nil {
		log.Println("重新载入失败", err)
	}
}