package str

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"testing"
)

func BenchmarkAdd(t *testing.B) {
	a := "hello"
	b := "wolrd"
	for i := 0; i < t.N; i++ {
		_ = a + "," + b
	}
}

func BenchmarkAppend(t *testing.B) {
	a := "hello"
	b := "wolrd"
	c := make([]byte, 0, len(a)+len(b)+1)
	for i := 0; i < t.N; i++ {
		_ = string(append(append(append(c, a...), ','), b...))
	}
}

func BenchmarkFmt(t *testing.B) {
	a := "hello"
	b := "wolrd"
	for i := 0; i < t.N; i++ {
		_ = fmt.Sprintf("%s,%s", a, b)
	}
}

func BenchmarkJoin(t *testing.B) {
	a := "hello"
	b := "wolrd"
	for i := 0; i < t.N; i++ {
		_ = strings.Join([]string{a, b}, ",")
	}
}

func BenchmarkWrite(t *testing.B) {
	a := "hello"
	b := "wolrd"
	for i := 0; i < t.N; i++ {
		var buf strings.Builder
		buf.WriteString(a)
		buf.WriteString(",")
		buf.WriteString(b)
		_ = buf.String()
	}
}

type Student struct {
	Name   string
	Age    int32
	Remark [1024]byte
}

var buf, _ = json.Marshal(Student{Name: "Tom", Age: 19})
var pol = sync.Pool{
	New: func() any {
		return new(Student)
	},
}

func BenchmarkUnmarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stu := &Student{}
		_ = json.Unmarshal(buf, stu)
	}

}

func BenchmarkUnmarshalWithPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stu := pol.Get().(*Student)
		json.Unmarshal(buf, stu)
		pol.Put(stu)
	}
}
