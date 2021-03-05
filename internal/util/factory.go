package util

type GenericFactory struct {
	factoryMap map[string]interface{}
}

func InitFactory() *GenericFactory {
	return &GenericFactory{
		factoryMap: make(map[string]interface{}),
	}
}

func (factory *GenericFactory) Register(name string, object interface{})  {
	factory.factoryMap[name] = object
}

func (factory *GenericFactory) Get(name string) interface{} {
	return factory.factoryMap[name]
}

