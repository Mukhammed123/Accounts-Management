package numbers

import (
// "constraint"
)

// FIXME: 等Go 1.18發布後使用泛型

func IntMax(v0 int, vs ...int) int {
	if len(vs) == 0 {
		return v0
	}
	result := v0
	for _, v := range vs {
		if result < v {
			result = v
		}
	}
	return result
}

// func Max[V constraint.Ordered](v0 V, vs ...V) bool {
// 	if len(vs) == 0 {
// 		return v0
// 	}
// 	result := v0
// 	for _, v := range vs {
// 		if result < v {
// 			result = v
// 		}
// 	}
// 	return result
// }

func IntMin(v0 int, vs ...int) int {
	if len(vs) == 0 {
		return v0
	}
	result := v0
	for _, v := range vs {
		if result > v {
			result = v
		}
	}
	return result
}

// func IntMin[V constraint.Ordered](v0 V, vs ...V) bool {
// 	if len(vs) == 0 {
// 		return v0
// 	}
// 	result := v0
// 	for _, v := range vs {
// 		if result > v {
// 			result = v
// 		}
// 	}
// 	return result
// }
