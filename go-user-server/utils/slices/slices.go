package slices

import (
	// "constraint"
	"fmt"
	"reflect"
)

// FIXME: 等Go 1.18發布後使用泛型

func Contain(s interface{}, v interface{}) bool {
	if reflect.TypeOf(s).Kind() != reflect.Slice {
		panic(fmt.Errorf("slices.Contain: %v is not a slice", s))
	}
	sv := reflect.ValueOf(s)
	for i := 0; i < sv.Len(); i += 1 {
		if sv.Index(i).Interface() == v {
			return true
		}
	}
	return false
}

// func Contain[V constraint.Ordered](s []V, v V) bool {
// 	for _, value := range s {
// 		if value == v {
// 			return true
// 		}
// 	}
// 	return false
// }
