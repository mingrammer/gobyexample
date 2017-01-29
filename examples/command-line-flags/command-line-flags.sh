# 커맨드라인 플래그 프로그램을 실험하려면 우선 이를 컴파일한 후 직접 결과 바이너리를 실행하는 것이 가장 좋습니다.
$ go build command-line-flags.go

# 처음엔 모든 플래그에 값을 주어 프로그램을 실행해 봅시다.
$ ./command-line-flags -word=opt -numb=7 -fork -svar=flag
word: opt
numb: 7
fork: true
svar: flag
tail: []

# 참고로 생략된 플래그들은 자동으로 기본값을 갖게 됩니다.
$ ./command-line-flags -word=opt
word: opt
numb: 42
fork: false
svar: bar
tail: []

# 후미에 위치한 인자들은 모든 플래그 뒤에 위치할 수 있습니다.
$ ./command-line-flags -word=opt a1 a2 a3
word: opt
...
tail: [a1 a2 a3]

# 참고로 `flag` 패키지에서 모든 플래그들은 위치 인자 전에 나타나야합니다. (그렇지 않으면 플래그들은 위치 인자로 해석됩니다)
$ ./command-line-flags -word=opt a1 a2 a3 -numb=7
word: opt
numb: 42
fork: false
svar: bar
tail: [a1 a2 a3 -numb=7]

# `-h` 또는 `--help` 플래그를 사용하면 커맨드라인 프로그램에 대해 자동으로 생성된 도움말을 볼 수 있습니다. 
$ ./command-line-flags -h
Usage of ./command-line-flags:
  -fork=false: a bool
  -numb=42: an int
  -svar="bar": a string var
  -word="foo": a string

# `flag` 패키지로 지정되지 않은 플래그를 사용하면, 프로그램은 에러 메시지를 출력하며 도움말을 다시 보여줍니다.
$ ./command-line-flags -wat
flag provided but not defined: -wat
Usage of ./command-line-flags:
...

# 다음장에선 또 다른 프로그램을 매개변수화하는 일반적인 방법인 환경변수에 대해 살펴보겠습니다.