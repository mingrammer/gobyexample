# 프로그램을 실행시키면 프로그램은 `ls`로 대체됩니다.
$ go run execing-processes.go
total 16
drwxr-xr-x  4 mark 136B Oct 3 16:29 .
drwxr-xr-x 91 mark 3.0K Oct 3 12:50 ..
-rw-r--r--  1 mark 1.3K Oct 3 16:28 execing-processes.go

# Go가 고전 Unix의 `fork`를 제공하지 않는다는 사실을 기억하십시오.
#  보통 고루틴을 이용하여 프로세스를 생성하거나 `exec`하여
#  `fork`의 대부분의 use case를 다루기 때문에
#  웬만해선 문제가 되지 않을 것입니다.
