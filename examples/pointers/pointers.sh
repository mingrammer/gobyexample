# `zeroval`은 `main`에 있는 `i`값을 바꾸지 않지만,
# `zeroptr`은 해당 값의 메모리 주소를 참조하고 있기 때문에 값을 바꿀 수 있습니다.
$ go run pointers.go
initial: 1
zeroval: 1
zeroptr: 0
pointer: 0x42131100
