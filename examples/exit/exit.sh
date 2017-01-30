# 만약 `exit.go`를 `go run`을 이용하여 실행시키면,
#  exit 값을 `go`가 가져가서 출력될 것입니다.
$ go run exit.go
exit status 3

# 바이너리를 빌드하여 실행시키는 경우, 터미널에서 
#  status를 볼 수 있습니다.
$ go build exit.go
$ ./exit
$ echo $?
3

# `!`는 절대로 출력되지 않는다는 것을 기억하세요!
