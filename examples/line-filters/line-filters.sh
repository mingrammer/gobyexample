# 라인 필터를 실행해 보기위해, 몇 개의 소문자 라인들을 가진 파일을 생성합니다.
$ echo 'hello'   > /tmp/lines
$ echo 'filter' >> /tmp/lines

# 다음으로 대문자 라인들을 얻기 위해 라인 필터를 사용합니다.
$ cat /tmp/lines | go run line-filters.go
HELLO
FILTER
