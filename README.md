# Go Study

Go를 공부하며 작성한 예제를 모아 둔 학습 프로젝트입니다. 현재는 고루틴과 동시성 프로그래밍의 기초를 다루는 `day01` 예제가 있습니다.

## 프로젝트 구조

```text
.
├── go.mod
└── day01/
    ├── main.go
    ├── goroutine_waitgroup.go
    ├── channel_unbuffered.go
    ├── channel_buffered.go
    ├── channel_exercise_2_1.go
    ├── channel_exercise_2_2.go
    ├── channel_exercise_2_3.go
    ├── channel_exercise_2_4.go
    ├── context.go
    └── mutex.go
```

각 Day 폴더의 `main.go`가 해당 학습 예제의 실행 진입점입니다. `day01`에서는 다음 주제를 다룹니다.

- 고루틴과 `sync.WaitGroup`
- 버퍼가 있는 채널과 버퍼가 없는 채널
- 채널을 활용하는 연습 문제
- `context`를 이용한 작업 취소와 타임아웃
- `sync.Mutex`를 이용한 공유 데이터 보호

## 실행 방법

프로젝트 루트에서 실행합니다.

```sh
go run ./day01
```

현재 `day01/main.go`는 뮤텍스 예제를 실행합니다. 다른 예제를 실행하려면 `main.go`에서 원하는 함수 호출의 주석을 해제하고, 실행하지 않을 함수 호출은 주석 처리하세요.

## 테스트

모든 패키지를 확인하려면 프로젝트 루트에서 실행합니다.

```sh
go test ./...
```
