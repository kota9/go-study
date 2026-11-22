package main

import "fmt"

// 숫자를 Channel로 전달하는 함수
func sendExercise2_3Values(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i // 모든 숫자를 Channel에 전달
	}

	close(ch) // Channel 닫기
}

func runExercise2_3() {
	ch := make(chan int) // Channel 생성

	go sendExercise2_3Values(ch) // Goroutine 실행

	// 모든 데이터를 수신해 합계를 계산
	sum := 0
	for value := range ch {
		sum += value
	}
	fmt.Println("합계:", sum)
}
