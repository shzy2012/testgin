package bench

import (
	"net/http"
	"testing"

	"github.com/shzy2012/testgin/httputil"
)

var client *http.Client
var url string

func init() {

	client = httputil.NewHttpClient(100, 100, 3)
	url = "http://localhost:8080/v1/"
}
func Benchmark_Gin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		client.Get(url)
	}
}
