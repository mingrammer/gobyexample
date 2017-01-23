# 프로그램을 실행하면 패닉을 일으키며, 에러 메시지와 고루틴 트레이스 정보, 그리고 0이 아닌 상태값과 함께 종료됩니다.
$ go run panic.go
panic: a problem

goroutine 1 [running]:
main.main()
	/.../panic.go:12 +0x47
...
exit status 2

# 참고로 에러를 핸들링 하기 위해 익셉션을 사용하는 타 언어와는 달리, Go에서는 에러 가능성이 있는 곳에서 에러를 가리키는 반환값을 사용하는게 일반적입니다.
