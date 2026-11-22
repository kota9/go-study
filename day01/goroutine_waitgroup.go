package main

import (
	"fmt"
	"sync"
	"time"
)

// 각각의 작업을 실행하는 함수
func waitGroupWorker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d 시작\n", id)

	time.Sleep(1 * time.Second)

	fmt.Printf("Worker %d 완료\n", id)
}

func runGoroutineWaitGroup() {
	var wg sync.WaitGroup

	// 3개의 Goroutine 실행
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go waitGroupWorker(i, &wg)
	}

	// 모든 작업이 완료될 때까지 대기
	wg.Wait()

	fmt.Printf("모든 작업 종료")
}
