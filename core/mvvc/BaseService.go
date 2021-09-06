package mvvc

type serviceInterface interface {
	Filter() *serviceInterface
}

type BaseServiceStruct struct {
	serviceInterface
	Controller ControllerInterface
}

func (that *BaseServiceStruct) Filter() *BaseServiceStruct {
	return that
}