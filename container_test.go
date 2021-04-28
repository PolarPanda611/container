package container

import (
	"reflect"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

type test1 struct{}
type test2 struct{}

var (
	pool1 = &sync.Pool{New: func() interface{} { return &test1{} }}
	pool2 = &sync.Pool{New: func() interface{} { return &test2{} }}
)

func TestNewContainer(t *testing.T) {
	c := Config{
		AutoFree:  true,
		AutoWired: false,
	}
	tests := []struct {
		name string
		c    []Config
		want *Container
	}{
		// TODO: Add test cases.
		{
			name: "1",
			c:    []Config{},
			want: &Container{
				c:               DefaultConfig,
				poolMap:         make(map[reflect.Type]*sync.Pool),
				poolTags:        make(map[string]reflect.Type),
				instanceMapping: make(map[string]reflect.Type),
			},
		},
		{
			name: "2",
			c:    []Config{c},
			want: &Container{
				c:               &c,
				poolMap:         make(map[reflect.Type]*sync.Pool),
				poolTags:        make(map[string]reflect.Type),
				instanceMapping: make(map[string]reflect.Type),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewContainer(tt.c...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewContainer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainer_NewInstance(t *testing.T) {

	type fields struct {
		poolMap          map[reflect.Type]*sync.Pool
		instanceTypeList []reflect.Type
		poolTags         map[string]reflect.Type
		instanceMapping  map[string]reflect.Type
	}
	type args struct {
		instanceType reflect.Type
		instancePool *sync.Pool
		instanceTag  []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{
			name: "1",
			fields: fields{
				poolMap: map[reflect.Type]*sync.Pool{
					reflect.TypeOf(&test1{}): pool1,
				},
				instanceTypeList: []reflect.Type{reflect.TypeOf(&test1{})},
				poolTags: map[string]reflect.Type{
					"test1": reflect.TypeOf(&test1{}),
				},
				instanceMapping: map[string]reflect.Type{},
			},
			args: args{
				instanceType: reflect.TypeOf(&test1{}),
				instancePool: &sync.Pool{New: func() interface{} { return &test1{} }},
				instanceTag:  []string{"test1"},
			},
		},
		{
			name: "2",
			fields: fields{
				poolMap:          map[reflect.Type]*sync.Pool{},
				instanceTypeList: []reflect.Type{},
				poolTags:         map[string]reflect.Type{},
				instanceMapping:  map[string]reflect.Type{},
			},
			args: args{
				instanceType: reflect.TypeOf(&test1{}),
				instancePool: pool1,
				instanceTag:  []string{"test1"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Container{
				poolMap:          tt.fields.poolMap,
				instanceTypeList: tt.fields.instanceTypeList,
				poolTags:         tt.fields.poolTags,
				instanceMapping:  tt.fields.instanceMapping,
			}
			s.NewInstance(tt.args.instanceType, tt.args.instancePool, tt.args.instanceTag...)
			assert.Equal(t, &Container{
				poolMap: map[reflect.Type]*sync.Pool{
					reflect.TypeOf(&test1{}): pool1,
				},
				instanceTypeList: []reflect.Type{reflect.TypeOf(&test1{})},
				poolTags: map[string]reflect.Type{
					"test1": reflect.TypeOf(&test1{}),
				},
				instanceMapping: map[string]reflect.Type{},
			}, s, "instance not register correctly")
		})
	}
}

func TestContainer_GetInstanceTypeByTag(t *testing.T) {
	p1 := NewContainer()
	p1.NewInstance(reflect.TypeOf(&test1{}), pool1, "test1")
	p1.NewInstance(reflect.TypeOf(&test1{}), pool1, "test2")
	p1.NewInstance(reflect.TypeOf(&test2{}), pool2, "test3")
	p1.NewInstance(reflect.TypeOf(&test2{}), pool2, "test4")
	type fields struct {
		poolMap          map[reflect.Type]*sync.Pool
		instanceTypeList []reflect.Type
		poolTags         map[string]reflect.Type
		instanceMapping  map[string]reflect.Type
	}
	type args struct {
		tags []string
	}
	tests := []struct {
		name   string
		fields *Container
		args   args
		want   []reflect.Type
	}{
		// TODO: Add test cases.
		{
			name:   "1",
			fields: p1,
			args: args{
				tags: []string{"test1"},
			},
			want: []reflect.Type{reflect.TypeOf(&test1{})},
		},
		{
			name:   "2",
			fields: p1,
			args: args{
				tags: []string{},
			},
			want: []reflect.Type{reflect.TypeOf(&test1{}), reflect.TypeOf(&test2{})},
		},
		{
			name:   "3",
			fields: p1,
			args: args{
				tags: []string{"test1", "test3"},
			},
			want: []reflect.Type{reflect.TypeOf(&test1{}), reflect.TypeOf(&test2{})},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.fields.GetInstanceTypeByTag(tt.args.tags...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Container.GetInstanceTypeByTag() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContainer_CheckInstanceNameIfExist(t *testing.T) {
	p1 := NewContainer()
	p1.NewInstance(reflect.TypeOf(&test1{}), pool1, "test1")
	type args struct {
		instanceName reflect.Type
	}
	tests := []struct {
		name   string
		fields *Container
		args   args
		want   bool
	}{
		{
			name:   "1",
			fields: p1,
			args: args{
				instanceName: reflect.TypeOf(&test1{}),
			},
			want: true,
		},
		{
			name:   "2",
			fields: p1,
			args: args{
				instanceName: reflect.TypeOf(&test2{}),
			},
			want: false,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Container{
				poolMap:          tt.fields.poolMap,
				instanceTypeList: tt.fields.instanceTypeList,
				poolTags:         tt.fields.poolTags,
				instanceMapping:  tt.fields.instanceMapping,
			}
			if got := s.CheckInstanceNameIfExist(tt.args.instanceName); got != tt.want {
				t.Errorf("Container.CheckInstanceNameIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
