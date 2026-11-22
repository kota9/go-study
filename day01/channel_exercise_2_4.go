package main

import "fmt"

func runExercise2_4() {
	ch := make(chan int, 5) // 버퍼가 있는 Channel 생성

	for i := 1; i <= 5; i++ {
		ch <- i * 10 // 숫자를 Channel에 전달
	}

	close(ch) // Channel 닫기

	// Channel에서 숫자를 받아 출력
	for value := range ch {
		fmt.Println(value)
	}

	fmt.Println("모든 데이터 수신 완료")
}
