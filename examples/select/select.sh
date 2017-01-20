# 예상대로 `"one"`을 받은 후 `"two"`를 받습니다.
$ time go run select.go 
received one
received two

# 참고로 총 실행 시간은 2초 정도밖에 안되는데 1초와 2초 `Sleeps`이 동시에 실행되기 때문입니다.
real	0m2.245s
