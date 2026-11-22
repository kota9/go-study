package main

import (
	"fmt"
	"sync"
)

// Mutex를 이용한 공유 데이터 보호
func runMutex() {
	count := 0

	var mu sync.Mutex // Mutex 생성
	var wg sync.WaitGroup

	// 100개의 Goroutine 생성
	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			// 각 Goroutine에서 100번 증가
			for j := 0; j < 100; j++ {
				mu.Lock() // Mutex 잠금

				count++ // 공유 데이터 증가

				mu.Unlock() // Mutex 잠금 해제
			}
		}()
	}

	// 모든 Goroutine 종료 대기
	wg.Wait()

	fmt.Printf("최종 count: %d\n", count)
}
