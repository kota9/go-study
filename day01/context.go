package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Context를 이용한 작업 취소 및 Timeout
func runContext() {
	var wg sync.WaitGroup

	// --------------------------------
	// 1. WithCancel 실습
	// --------------------------------
	fmt.Println("=== WithCancel 실습 ===")

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)

	go contextWorker(ctx, &wg, "WithCancel")

	// 1.5초 후에 작업 취소
	time.Sleep(1500 * time.Millisecond)
	cancel() // 작업 취소

	// 작업 종료 대기
	wg.Wait()

	fmt.Println("WithCancel 작업 종료")

	// --------------------------------
	// 2. WithTimeout 실습
	// --------------------------------
	fmt.Println()
	fmt.Println("=== WithTimeout 실습 ===")

	ctx, cancel = context.WithTimeout(
		context.Background(),
		1500*time.Millisecond,
	)
	defer cancel() // 작업 종료 시 cancel 호출

	wg.Add(1)

	go contextWorker(ctx, &wg, "WithTimeout")

	// Timeout 또는 작업 종료 대기
	wg.Wait()

	fmt.Println("WithTimeout 작업 종료")
}

// Context를 확인하면서 실행하는 작업
func contextWorker(ctx context.Context, wg *sync.WaitGroup, name string) {
	defer wg.Done()

	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("[%s] 작업 취소됨: %v\n", name, ctx.Err())
			return

		case t := <-ticker.C:
			fmt.Printf("[%s] 작업 진행 중: %v\n", name, t)
		}
	}
}
