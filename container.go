package container

import (
	"reflect"
	"sync"
)

type Config struct {
	AutoFree  bool
	AutoWired bool
}

var (
	DefaultConfig = &Config{
		AutoFree:  true,
		AutoWired: true,
	}
)

type Container struct {
	c *Config
	// map[instanceName] = instancePool
	poolMap map[reflect.Type]*sync.Pool
	// instance list
	instanceTypeList []reflect.Type
	// instance tags maps instanceName
	poolTags map[string]reflect.Type
	// instanceMapping instance mapping
	// caching the instance di instance relation during the self check
	instanceMapping map[string]reflect.Type
}

// NewContainer new pool with init map
func NewContainer(c ...Config) *Container {
	result := new(Container)
	result.poolMap = make(map[reflect.Type]*sync.Pool)
	result.poolTags = make(map[string]reflect.Type)
	result.instanceMapping = make(map[string]reflect.Type)
	if len(c) > 0 {
		result.c = &c[0]
	} else {
		result.c = DefaultConfig
	}
	return result
}

// NewInstance add new instance
func (s *Container) NewInstance(instanceType reflect.Type, instancePool *sync.Pool, instanceTag ...string) {
	if _, ok := s.poolMap[instanceType]; ok {
		return
	}
	s.poolMap[instanceType] = instancePool
	s.instanceTypeList = append(s.instanceTypeList, instanceType)
	if len(instanceTag) > 0 {
		if instanceTag[0] != "" {
			s.poolTags[instanceTag[0]] = instanceType
		}
	}
}

// GetInstanceType get all service type by tag
// if no tag provide , return all type
// if tags provide , will return the types of the tags
func (s *Container) GetInstanceTypeByTag(tags ...string) []reflect.Type {
	if len(tags) > 0 {
		var types []reflect.Type
		for _, v := range tags {
			if instance, ok := s.poolTags[v]; ok {
				types = append(types, instance)
			}
		}
		return types
	}
	return s.instanceTypeList
}

// CheckInstanceNameIfExist check contain name if exist
func (s *Container) CheckInstanceNameIfExist(instanceName reflect.Type) bool {
	_, ok := s.poolMap[instanceName]
	return ok
}

// InstanceMapping get instance mapping
// return the copy of the instance mapping
func (s *Container) InstanceMapping() map[string]reflect.Type {
	instanceMap := make(map[string]reflect.Type, len(s.instanceMapping))
	for k, v := range s.instanceMapping {
		instanceMap[k] = v
	}
	return instanceMap
}

func (s *Container) GetInstance(instanceType reflect.Type, injectingMap map[reflect.Type]interface{}) (interface{}, map[reflect.Type]interface{}, map[reflect.Type]interface{}) {
	return nil, nil, nil
}
