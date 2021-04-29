package container

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type inject1 struct {
	name string
}
type testInject struct {
	test1 inject1  `container:"autowire:true;autofree:true"`
	Test2 inject1  `container:"autowire:true;autofree:true"`
	test3 inject1  `container:"autowire:true;autofree:true"`
	test4 inject1  `container:"autowire:true;autofree:true"`
	Test5 *inject1 `container:"autowire:true;autofree:true"`
}

func TestContainer_DiFree(t *testing.T) {
	a := inject1{"1"}
	empty := inject1{}
	{
		obj := struct{}{}
		assert.PanicsWithError(t, "toFree object struct {} should be addressable", func() { NewContainer().DiFree(obj) }, "addressable")
	}
	{
		obj := struct {
			test1 inject1 `container:"autowire:true;autofree:true"`
		}{a}
		NewContainer().DiFree(&obj)
		assert.Equal(t, a, obj.test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 inject1 `container:"autowire:true;autofree:true"`
		}{a}
		NewContainer().DiFree(&obj)
		assert.Equal(t, empty, obj.Test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 inject1 `container:"autowire:false;autofree:true"`
		}{a}
		NewContainer().DiFree(&obj)
		assert.Equal(t, a, obj.Test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 inject1 `container:"autowire:true;autofree:false"`
		}{a}
		NewContainer().DiFree(&obj)
		assert.Equal(t, a, obj.Test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 inject1 `container:"autowire:false;autofree:false"`
		}{a}
		NewContainer().DiFree(&obj)
		assert.Equal(t, a, obj.Test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 *inject1 `container:"autowire:true;autofree:true"`
		}{&a}
		NewContainer().DiFree(&obj)
		assert.Nil(t, obj.Test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 *inject1 `container:"autowire:false;autofree:true"`
		}{&a}
		NewContainer().DiFree(&obj)
		assert.Equal(t, &a, obj.Test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 *inject1 `container:"autowire:true;autofree:false"`
		}{&a}
		NewContainer().DiFree(&obj)
		assert.Equal(t, &a, obj.Test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 *inject1 `container:"autowire:false;autofree:false"`
		}{&a}
		NewContainer().DiFree(&obj)
		assert.Equal(t, &a, obj.Test1, "diFree error ")
	}
	{
		obj := struct {
			Test1 string `container:"autowire:true;autofree:true"`
		}{"1"}
		NewContainer().DiFree(&obj)
		assert.Equal(t, "", obj.Test1, "diFree error ")
	}
	{
		s := "1"
		obj := struct {
			Test1 *string `container:"autowire:true;autofree:true"`
		}{&s}
		NewContainer().DiFree(&obj)
		assert.Nil(t, obj.Test1, "diFree error ")
	}
	{
		s := "1"
		obj := struct {
			Test1 *string `container:"autowire:true;autofree:true"`
			T     string
		}{&s, "1"}
		NewContainer().DiFree(&obj)
		assert.Nil(t, obj.Test1, "diFree error ")
		assert.Equal(t, "1", obj.T, "diFree error ")
	}
}
