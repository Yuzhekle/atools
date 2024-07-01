package bench

import (
	"testing"
)

// func main() {
// 	fmt.Println("num: 1==1", num(1, 1))
// 	fmt.Println("bit: 1==1", bit(1, 1))
// }

func num(a, b int) bool {
	return a == b
}

func bit(a, b int) bool {
	return ^(a ^ b) == -1
}

func BenchmarkNum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		num(1, 1)
	}
}

func BenchmarkBit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bit(1, 1)
	}
}
