package sys

import "testing"

var int8ptr *c_char

func BenchmarkCString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		int8ptr = CString("s1 != s2")
	}
}
