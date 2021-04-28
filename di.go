package container

// func DiFree(dest interface{}) {
// 	destVal := reflect.Indirect(reflect.ValueOf(dest))
// 	for index := 0; index < destVal.NumField(); index++ {
// 		val := destVal.Field(index)
// 		if !getTagByName(dest, index, AUTOWIRED) {
// 			continue
// 		}
// 		if !GetAutoFreeTag(dest, index) {
// 			continue
// 		}
// 		if !val.CanSet() {
// 			continue
// 		}
// 		if !val.IsZero() {
// 			val.Set(reflect.Zero(val.Type()))
// 		}
// 	}
// }

// func GetAutoWiredTag(obj interface{}, index int) bool {

// }

// func GetAutoFreeTag(obj interface{}, index int) bool {
// 	v, exist := getTagByName(obj, index, AUTOFREE)
// 	if exist {

// 	}
// 	return true
// }
