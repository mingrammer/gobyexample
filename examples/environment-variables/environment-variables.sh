# 프로그램을 돌려보면
#  `FOO`의 경우 우리가 프로그램 안에서 설정한 값을 가져오는 반면,
#  `BAR`의 경우는 빈 문자열을 가져오는 것을 볼 수 있습니다.
$ go run environment-variables.go
FOO: 1
BAR: 

# 환경변수로 지정된 키의 리스트는 컴퓨터에 따라 다르게 보일 수 있습니다.
TERM_PROGRAM
PATH
SHELL
...

# 우리가 환경변수로 `BAR`를 먼저 설정해두면,
#  실행될 프로그램이 해당 키에 설정된 값을 가져갑니다.
$ BAR=2 go run environment-variables.go
FOO: 1
BAR: 2
...
