package synctest

import (
	"golang.org/x/sync/singleflight"
	"sync/atomic"
	"testing"
)

var (
	count int64 = 0
)

func BenchmarkSingleFlight1(b *testing.B) {
	singleFlight := singleflight.Group{}
	for i:=0;i<b.N;i++{
		go func() {
			_ ,_,_= singleFlight.Do("testkey", func() (i interface{}, e error) {
				return countAddOne(), nil
			})
		}()
	}
}

func BenchmarkSingleFlight2(b *testing.B) {
	singleFlight := singleflight.Group{}
	for i:=0;i<b.N;i++{
		_ ,_,_= singleFlight.Do("testkey", func() (i interface{}, e error) {
			return countAddOne(), nil
		})
	}
}

func countAddOne() int64 {
	return atomic.AddInt64(&count, 1)
}