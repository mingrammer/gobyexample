// 채널 _Closing_은 더 이상 채널에 보낼 데이터가 없음을 나타냅니다.
//  이는 채널의 리시버들에게 완료 상태를 전달하는데에 유용합니다.

package main

import "fmt"

// 이 예제에서 우리는 `main()` 고루틴에서 워커 고루틴으로 작업을 전달하기 위해 `jobs` 채널을 사용합니다.
//  워커에서 돌릴 잡이 더 이상 없을 경우, `jobs` 채널을 `close` 합니다.
func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	// 여기에 워커 고루틴 하나가 있습니다. 이는 `j, more := <-jobs`을 통해 `jobs`로부터 반복적으로 값을 수신합니다.
	//  이 두 값을 반환하는 특별한 형태의 수신값에서, `more`값은 `jobs`가 `close`되고 채널에 있는 모든 값들이 수신될 경우 `false`값을 갖게됩니다.
	//  모든 잡이 종료되었음을 알리기 위해 `done`을 사용합니다.
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// `jobs` 채널을 통해 워커로 3개의 잡을 보낸 후, 채널을 닫습니다.
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")

	// 이전에 봤던 [synchronization](/gobyexample/channel-synchronization) 방법으로 워커를 대기합니다.
	<-done
}
