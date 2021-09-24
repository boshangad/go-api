package v1

import (
	"github.com/boshangad/go-api/api/v1/public"
)

type ApiGroup struct {
	Public public.ApiGroup
}

var ApiGroupApp = new(ApiGroup)