package v1

import "github.com/boshangad/go-api/routers/v1/public"

type GroupApi struct {
	public.RouterApi
}

var ApiGroup = new(GroupApi)