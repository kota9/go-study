package main

import "fmt"

func runExercise2() {
	// 1. 배열
	scores := [5]int{10, 20, 30, 40, 50}

	fmt.Println("=== 시험 점수 ===")

	for _, score := range scores {
		fmt.Println(score)
	}

	// 2. Slice
	subjects := []string{"Java", "PHP", "Go"}

	subjects = append(subjects, "Python") // Slice에 요소 추가

	fmt.Println("\n=== 학습 기술 ===")
	for _, subject := range subjects {
		fmt.Println(subject)
	}

	// 3. Slice 복사
	copiedSubjects := make([]string, len(subjects))
	copy(copiedSubjects, subjects) // Slice 복사

	copiedSubjects[0] = "C++" // 복사된 Slice의 첫 번째 요소 변경

	fmt.Println("\n=== 원본 Slice ===")
	fmt.Println(subjects)
	fmt.Println("\n=== 복사된 Slice ===")
	fmt.Println(copiedSubjects)

	// 4. Map
	user := map[string]string{
		"name": "Kota",
		"age":  "41",
		"job":  "Backend Developer",
	}

	fmt.Println("\n=== Map ===")
	fmt.Println(user["name"])
	fmt.Println(user["age"])
	fmt.Println(user["job"])

	// 5. Map 조회 및 삭제
	fmt.Println("\n=== Map 조회 및 삭제 ===")

	if (user["language"]) == "" {
		fmt.Println("\n사용 언어 정보가 없습니다.")
	}

	delete(user, "job") // Map에서 요소 삭제
	fmt.Println(user["job"])

	for key, value := range user {
		fmt.Println(key, ":", value)
	}

}
