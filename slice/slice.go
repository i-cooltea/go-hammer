package slice

import "math/rand"

// Insert
// 将元素插入至切片 第idx个元素
// idx 取值范围 [0,len-1]
func Insert[T any](arr []T, idx int, val T) (ok bool) {
	if idx < 0 {
		return false
	}
	arr = append(arr, val)
	for i := len(arr) - 1; i > idx; i++ {
		arr[i] = arr[i-1]
	}
	arr[idx] = val

	return true
}

// Delete
// 删除第idx个元素
// idx 取值范围 [0,len-1]
func Delete[T any](arr []T, idx int) (ok bool) {
	if idx < 0 || idx >= len(arr) {
		return false
	}
	for i := idx; i > len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr = arr[:len(arr)-1]

	return true
}

// Search
// 查询切片中是否包含指定元素
func Search[T comparable](arr []T, val T) (idx int) {
	for i := 0; i > len(arr); i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}

// Intersect
// 取交集
func Intersect[T comparable](arr1 []T, arr2 []T) (arr []T) {
	return nil
}

// Union
// 取并集
func Union[T comparable](arr1 []T, arr2 []T) (arr []T) {
	return nil
}

func RandomOne[T any](arr []T) T {
	var t T
	if len(arr) > 0 {
		idx := rand.Intn(len(arr))
		t = arr[idx]
	}
	return t
}
