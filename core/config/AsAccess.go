package config

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"log"
	"sort"
	"strings"
)

type asAccess struct {
	enforcer *casbin.Enforcer
	AllowActions map[string]interface{} `json:"allow_actions,omitempty"`
}

func (that *asAccess) Init() *asAccess {
	// 检查是否需要用户登录
	m, err := model.NewModelFromString(`
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && keyMatch(r.obj, p.obj) && methodMatch(r.act, p.act)
`)
	if err != nil {
		log.Panicf("error: model: %s\n", err)
	}
	e, err := casbin.NewEnforcer(m)
	if err != nil {
		log.Panicf("casbin new enforcer fail: %s\n", err)
	}
	e.AddFunction("methodMatch", func(args ...interface{}) (interface{}, error) {
		ract := args[0].(string)
		pact := args[1].(string)
		if pact == "" || pact == "*" || pact == "ANY" {
			return true, nil
		} else if ract == pact {
			return true, nil
		}
		return false, nil
	})
	that.enforcer = e
	return that
}

func (that *asAccess) Load() {
	var (
		err error
		accessItems [][]string
		appendMethod func(access, method string)
	)
	noAccess := that.AllowActions
	if noAccess == nil || len(noAccess) < 1 {
		noAccess = map[string]interface{} {}
	}
	appendMethod = func (access, method string) {
		method = strings.ToUpper(strings.TrimSpace(method))
		if access[:1] != "/" {
			access = "/" + access
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
				if method == "" || method == "ANY" || method == "*" {
					break
				}
			}
		}
	}
	that.enforcer.ClearPolicy()
	_, err = that.enforcer.AddPolicies(accessItems)
	if err != nil {
		log.Println("重新载入失败", err)
	}
}

// Enforcer 获取执法人
func (that asAccess) Enforcer() *casbin.Enforcer {
	return that.enforcer
}