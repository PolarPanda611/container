package container

import "reflect"

func getTagByName(object interface{}, index int, name Keyword) (string, bool) {
	objectType := reflect.TypeOf(object)
	switch objectType.Kind() {
	case reflect.Struct:
		return objectType.Field(index).Tag.Lookup(string(name))
	case reflect.Ptr:
		return objectType.Elem().Field(index).Tag.Lookup(string(name))
	default:
		panic("wrong type , must be struct or ptr ")
	}
}
