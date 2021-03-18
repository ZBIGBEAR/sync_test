package synctest

import (
	"context"
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/semaphore"
	"math/rand"
	"testing"
	"time"
)


func doJob(){
	time.Sleep(time.Microsecond)
}
func BenchmarkTestNoSemaphore(b *testing.B) {
	for i:=0;i<b.N;i++{
		go func() {
			doJob()
		}()
	}
}

func BenchmarkTestNoSemaphoreNoGo(b *testing.B) {
	for i:=0;i<b.N;i++{
		doJob()
	}
}

func BenchmarkSemaphore1(b *testing.B) {
	sema := semaphore.NewWeighted(1)
	ctx := context.Background()
	for i:=0;i<b.N;i++{
		go func() {
			err := sema.Acquire(ctx, 1)
			assert.Equal(b, nil, err)
			doJob()
			sema.Release(1)
		}()
	}
}

func BenchmarkTestSemaphore4(b *testing.B) {
	sema := semaphore.NewWeighted(4)
	ctx := context.Background()
	for i:=0;i<b.N;i++{
		go func() {
			count := rand.Int63n(4)
			err := sema.Acquire(ctx, count)
			assert.Equal(b, nil, err)
			doJob()
			sema.Release(count)
		}()
	}
}

func BenchmarkTestSemaphore10(b *testing.B) {
	sema := semaphore.NewWeighted(10)
	ctx := context.Background()
	for i:=0;i<b.N;i++{
		go func() {
			count := rand.Int63n(10)
			err := sema.Acquire(ctx, count)
			assert.Equal(b, nil, err)
			doJob()
			sema.Release(count)
		}()
	}
}