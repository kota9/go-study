package main

import "fmt"

func runUserInfo() {
	// 1. 변수 선언
	name := "Kota"
	age := 41
	active := true

	fmt.Println("=== 사용자 정보 ===")
	fmt.Println("이름:", name)
	fmt.Println("나이:", age)
	fmt.Println("활성 상태:", active)

	// 2. 배열
	scores := [3]int{80, 90, 100}
	fmt.Println("\n=== 시험 점수 ===")

	for index, score := range scores {
		fmt.Println(index, score)
	}

	// 3. Slice
	subjects := []string{"Java", "PHP", "Go"}

	subjects = append(subjects, "Spring") // Slice에 요소 추가

	fmt.Println("\n=== 학습 기술 ===")

	for _, subject := range subjects {
		fmt.Println(subject)
	}

	// 4. Map
	user := map[string]string{
		"name": "Kota",
		"job":  "Backend Developer",
	}

	user["language"] = "Go" // Map에 요소 추가

	fmt.Println("\n=== 사용자 상세 정보 ===")

	for key, value := range user {
		fmt.Println(key, ":", value)
	}

	// 5. Map 데이터 조회
	language, exists := user["language"]

	if exists {
		fmt.Println("\n사용 언어:", language)
	}

	// 6. Map 데이터 삭제
	delete(user, "job") // Map에서 요소 삭제

	fmt.Println("\n삭제 후 사용자 정보:", user)
}
