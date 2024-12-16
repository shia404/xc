package main

import (
	"fmt"
	"github.com/shia404/xc/pkg"
	"sync"
	"testing"
)

func Test_main(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(pkg.Snowflake.Generate().Int64())
		}()
	}
	wg.Wait()
}
