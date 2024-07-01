package main

// import "contianer/heap"

// func main() {
// 	heap.Init()
// }

// // O(logn)
// func find_fengzhi(s []int) (index int) {
// 	length := len(s)
// 	if length == 0 {
// 		return -1
// 	}
// 	if length == 1 {
// 		return 0
// 	}
// 	if length == 2 {
// 		if s[0] > s[1] {
// 			return 0
// 		} else {
// 			return 1
// 		}
// 	}
// 	left := 0
// 	right := length - 1
// 	for left < right {
// 		mid := (left + right) / 2
// 		if s[mid] > s[mid+1] {
// 			right = mid
// 		} else {
// 			left = mid + 1
// 		}
// 	}
// 	return left
// }
