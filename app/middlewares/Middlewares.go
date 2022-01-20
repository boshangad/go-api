package middlewares

import "github.com/boshangad/v1/app/controller"

var middlewares = make(map[string]func(c *controller.Context), 2048)

func AddMiddleware(key string, fun func(c *controller.Context)) {
	if fun == nil {
		delete(middlewares, key)
		return
	}
	middlewares[key] = fun
}

func DeleteMiddleware(key string) {
	delete(middlewares, key)
}

func GetMiddleware(key string) func(c *controller.Context) {
	return middlewares[key]
}
