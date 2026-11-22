package main

import "fmt"

// 숫자를 Channel로 전달하는 함수
func sendExercise2_1Values(ch chan int) {
	for i := 1; i <= 10; i++ {
		ch <- i // 숫자를 Channel에 전달
	}

	close(ch) // Channel 닫기
}

func runExercise2_1() {
	ch := make(chan int) // Channel 생성

	go sendExercise2_1Values(ch) // Goroutine 실행

	// Channel에서 숫자를 받아 출력
	for value := range ch {
		fmt.Println(value)
	}

	fmt.Println("모든 데이터 수신 완료")
}
