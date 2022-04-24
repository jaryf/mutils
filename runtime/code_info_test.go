package runtime

import "testing"

func BenchmarkRunFuncInfo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunFuncInfo()
	}
}
