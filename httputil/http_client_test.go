package httputil

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

var client *http.Client
var queue chan int
var done chan bool

func init() {
	//开启100 http client
	client = NewHttpClient(100, 100, 3)
	//队列长度 20
	queue = make(chan int, 20)
	done = make(chan bool)
}

func Test_HttpClient(b *testing.T) {

	t1 := time.Now()
	workSize := 100
	for i := 0; i < workSize; i++ {
		go worker()
	}

	for i := 0; i < 1000; i++ {
		queue <- i
	}

	close(queue)

	count := 0
	for range done {
		count++
		if count == workSize {
			return
		}
	}

	elapsed := time.Since(t1)
	log.Println("App elapsed: ", elapsed)
}

func worker() {
	for {
		index, ok := <-queue
		if !ok {
			done <- true
			return
		}
		resp, err := client.Get("https://www.baidu.com")
		if err != nil {
			log.Println(err)
		}

		defer resp.Body.Close()

		//ioutil.ReadAll(resp.Body)
		io.Copy(ioutil.Discard, resp.Body)
		log.Println("request=>", index)
	}
}
