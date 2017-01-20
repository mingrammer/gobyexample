// _Timeouts_은 외부 리소스를 연결하거나 실행 시간을 제한해야하는 프로그램에서 중요합니다.
//  Go에서 타임아웃 구현은 채널과 `select`를 사용하면 쉽고 우아하게 할 수 있습니다.
package main

import "time"
import "fmt"

func main() {

	// 이번 예제에선, 2초 후에 `c1` 채널에 결괏값을 반환하는 외부 호출을 실행한다고 가정합니다.
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 다음은 `select`로 구현한 타임아웃입니다.
	//  `res := <-c1`은 결괏값을 대기하고 `<-Time.After`는 1초의 타임아웃 후에 전달되는 값을 대기합니다.
	//  `select`는 대기중인 첫번째 리시브로 진행되고 있기 때문에, 만약 이 연산이 허용된 1초 보다 더 오래걸릴 경우, 타임아웃이 발생합니다.
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	// 만약 타움아웃을 3초로 늘리면, `c2`로부터의 수신은 성공하고 결괏값을 출력할겁니다.
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
