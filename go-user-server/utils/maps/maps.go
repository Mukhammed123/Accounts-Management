package maps

import (
	// "constraint"
	"fmt"
	"reflect"
)

// FIXME: 等Go 1.18發布後使用泛型

func Contain(m interface{}, v interface{}) bool {
	if reflect.TypeOf(m).Kind() != reflect.Map {
		panic(fmt.Errorf("MapContain: %v is not a map", m))
	}
	mv := reflect.ValueOf(m)
	for _, k := range mv.MapKeys() {
		if mv.MapIndex(k).Interface() == v {
			return true
		}
	}
	return false
}

// func Contain[K comparable, V constraint.Ordered](m map[K]V, v V) bool {
// 	for _, value := range m {
// 		if value == v {
// 			return true
// 		}
// 	}
// 	return false
// }

func KeyOf(m interface{}, v interface{}) interface{} {
	if reflect.TypeOf(m).Kind() != reflect.Map {
		panic(fmt.Errorf("maps.KeyOf: %v is not a map", m))
	}
	mv := reflect.ValueOf(m)
	for _, k := range mv.MapKeys() {
		if mv.MapIndex(k).Interface() == v {
			return k.Interface()
		}
	}
	return nil
}

// func KeyOf[K comparable, V constraint.Ordered](m map[K]V, v V) *K {
// 	for key, value := range m {
// 		if value == v {
// 			return &key
// 		}
// 	}
// 	return nil
// }
