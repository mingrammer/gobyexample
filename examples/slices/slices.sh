# 참고로 슬라이스는 배열과는 다른 타입이지만 `fmt.Println`로 출력을 하면 유사한 형태로 출력됩니다.
$ go run slices.go
emp: [  ]
set: [a b c]
get: c
len: 3
apd: [a b c d e f]
cpy: [a b c d e f]
sl1: [c d e]
sl2: [a b c d e]
sl3: [c d e f]
dcl: [g h i]
2d:  [[0] [1 2] [2 3 4]]

# Go 슬라이스의 설계와 구현체에 대한 자세한 정보는 Go팀에서 작성한 [great blog post](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html)를 참고하세요.

# 지금까지 배열과 슬라이스를 살펴봤습니다. 다음엔 Go의 또 다른 핵심 내장 데이터 구조인 maps를 살펴봅시다.
