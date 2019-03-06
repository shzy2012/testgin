package logger

import "testing"

var lg ILogger

func init() {
	lg = NewLogger()
}

func Benchmark_Logger_Info(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lg.Info(i)
	}
}

// func Benchmark_Logger_Rrror(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		lg.Error(i)
// 	}
// }
