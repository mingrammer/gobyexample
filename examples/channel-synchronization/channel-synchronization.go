// 고루틴간의 실행을 동기화하기 위해 채널을 사용할 수 있습니다.
//  고루틴이 끝날때까지 대기하기 위해 블로킹 리시브를 사용하는 예제를 보겠습니다.

package main

import "fmt"
import "time"

// 다음은 고루틴에서 실행하기 위한 함수입니다.
//  `done` 채널은 이 함수의 작업이 끝났음을 다른 고루틴에게 알리기 위해 사용됩니다.
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 작업의 종료를 알리기 위해 값을 전달합니다.
	done <- true
}

func main() {

	// 알림을 위한 채널을 전달하면서 워커 고루틴을 실행합니다.
	done := make(chan bool, 1)
	go worker(done)

	// 채널은 워커로부터 알림을 받을 때까지 블로킹됩니다.
	<-done
}
