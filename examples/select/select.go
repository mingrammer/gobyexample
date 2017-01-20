// Go의 _select_는 다중 채널 연산들을 대기할 수 있도록 해줍니다.
//  select을 사용한 고루틴과 채널의 조합은 Go의 강력한 기능입니다.

package main

import "time"
import "fmt"

func main() {

	// 이 예제에서 우리는 두 채널에서 select를 사용하겠습니다.
	c1 := make(chan string)
	c2 := make(chan string)

	// 각 채널은 한 가지 예로 동시에 실행되는 고루틴에서의 RPC 연산의 실행을 블로킹하는 경우를 시뮬레이션 하기위해 일정 시간 후에 값을 받습니다.
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	// `select`를 사용하여 동시에 이 값들을 대기하며 도착하는대로 각 값을 출력합니다.
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
