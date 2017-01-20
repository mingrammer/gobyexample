$ go run channel-synchronization.go      
working...done                  

# 만약 `<- done` 줄을 지우면 `worker`가 실행되기도 전에 프로그램이 종료될 수 있습니다.