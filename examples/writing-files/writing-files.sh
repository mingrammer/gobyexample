# 파일 쓰기 코드를 실행해보세요.
$ go run writing-files.go 
wrote 5 bytes
wrote 7 bytes
wrote 9 bytes

# 그리고 작성된 파일의 내용을 확인하세요.
$ cat /tmp/dat1
hello
go
$ cat /tmp/dat2
some
writes
buffered

# 다음으로 우리가 방금 본 파일 I/O 아이디어를 `stdin`과 `stdout` 스트림에 적용해 보겠습니다.