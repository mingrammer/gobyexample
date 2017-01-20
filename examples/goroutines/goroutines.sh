# 이 프로그램을 실행하면, 블로킹된 호출(동기적으로 호출된 함수)의 결과가 먼저 출력된 후에, 인터리브된 두 개의 고루틴의 결괏값이 출력됩니다.
#  이 인터리빙은 Go 런타임에 의해 고루틴들이 동시에 실행되는 프로세스를 반영합니다.
$ go run goroutines.go
direct : 0
direct : 1
direct : 2
goroutine : 0
going
goroutine : 1
goroutine : 2
<enter>
done

# 다음으로, Go 프로그램에서 동시에 실행되는 고루틴들을 보완해주는 채널(channels)에 대해 알아보겠습니다.