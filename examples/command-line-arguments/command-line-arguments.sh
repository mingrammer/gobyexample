# 커맨드라인 인자를 실험하려면 우선 `go build`로 바이너리를 빌드하는 것이 가장 좋습니다.
$ go build command-line-arguments.go
$ ./command-line-arguments a b c d
[./command-line-arguments a b c d]       
[a b c d]
c

# 다음으론 플래그를 활용한 고급 커맨드라인 처리에 대해 살펴보겠습니다.