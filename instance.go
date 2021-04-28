package container

import (
	"reflect"
	"sync"
)

type Instance interface {
	Interface() interface{}
	Release()
}
type instance struct {
	obj      interface{}
	resource map[reflect.Type]interface{}
}

func NewInstance(p *sync.Pool) Instance {

	return &instance{
		obj:      p.New(),
		resource: make(map[reflect.Type]interface{}),
	}

}
func (i *instance) getAutoFreeTag(index int) bool {
	return true
}

func (i *instance) diAllfields() {

}

func (i *instance) Interface() interface{} {

}

func (i *instance) Release() {

}

func (i *instance) diSelfCheck() {

}
