package synctest

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"sync/atomic"
	"testing"
)
func TestErrGroupCtx(t *testing.T) {
	var count int64 = 0
	eg, ctx := errgroup.WithContext(context.Background())
	for i:=0;i<10;i++{
		eg.Go(func() error {
			select {
			case <- ctx.Done():
				return nil
			default:
				atomic.AddInt64(&count,1)
				if count%2 == 0 {
					// time.Sleep(1000)
					return errors.New(fmt.Sprintf("err:%d", count))
				}else{
					fmt.Println("count:",count)
					return nil
				}
			}
		})
	}
	err := eg.Wait()
	// 这里的err有可能是2，4，6，8，10，是第一个报错的协程返回的err. count有可能是是1到10，因为每个协程在开始之前都
	// 通过ctx判断了是否已经取消了，如果有的协程报错了则已经取消了，当前协程不用运行。
	// 如果在22行sleep 1s,则count一定等于10.因为当count为基数，协程已经都指向完了，当count为偶数，协程都在sleep处
	// 等待，已经完成了count++的动作
	fmt.Println("err:",err,",count:",count)
}

func TestErrGroup(t *testing.T) {
	var count int64 = 0
	eg := errgroup.Group{}
	for i:=0;i<10;i++{
		eg.Go(func() error {
			atomic.AddInt64(&count,1)
			if count%2 == 0 {
				return errors.New(fmt.Sprintf("err:%d", count))
			}else{
				fmt.Println("count:",count)
				return nil
			}
		})
	}
	err := eg.Wait()
	// 这里的err有可能是2，4，6，8，10，count一定是10，因为每个go协程都会运行。但是err只会是第一个报错的协程
	fmt.Println("err:",err,",count:",count)
}