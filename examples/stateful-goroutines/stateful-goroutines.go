// 이전 예제에서 다중 고루틴의 공유 자원 접근을 동기화 하기위해 [mutexes](/gobyexample/mutexes)로 명시적인 잠금(locking)을 사용했습니다.
//  다른 방법인 고루틴과 채널의 내장 동기화 기능을 사용하여 같은 일을 할 수 있습니다.
//  채널 기반 접근법은 Go의 통신을 통한 메모리 공유와 정확히 한 고루틴이 각 데이터의 일부를 소유한다는 아이디어에 기반합니다.

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

// 이 예제에서 상태는 한 고루틴에 의해 소유됩니다. 이는 데이터가 동시 접근으로인해 손상되지 않음을 보장합니다.
//  상태를 읽거나 쓰기위해, 다른 고루틴들은 소유중인 고루틴으로 메시지를 보내고 응답을 받습니다.
//  다음 `readOps`와 `writeOps` 구조체는 요청들과 소유중인 고루틴이 응답하기 위한 방법을 캡슐화 합니다.
type readOp struct {
	key  int
	resp chan int
}
type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {

	// 이전과 같이 얼마나 연산이 얼마나 이루어지는지를 카운팅합니다.
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// `reads`와 `writes` 채널은 각각 읽기와 쓰기 요청을 전달하기 위함으로 다른 고루틴들에 의해 사용됩니다.
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// 다음 고루틴은 `state`를 소유하며, 이는 이전 예제에서처럼 map을 사용하지만 지금은 상태있는 고루틴이 이를 프라이빗으로 갖고 있습니다.
	//  이 고루틴은 반복적으로 `reads`와 `writes` 채널중 하나를 선택하며, 값이 도착하는대로 요청에 대한 응답 처리를 합니다.
	//  응답은 우선 요청된 연산을 수행함으로써 실행되며 `write`의 경우는 성공을 알리기 위해 (`reads`의 경우는 원하는 값을 ) 응답 채널 `resp`로 값을 전달합니다.
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 다음은 `reads` 채널을 통해 상태를 가진 고루틴에 읽기를 요청하기 위해 100개의 고루틴을 띄웁니다.
	//  각 읽기는 `readOp`을 생성해야하며, 이를 `reads` 채널을 통해 전달하고, `resp` 채널로 들어온 값을 통해 결괏값을 받습니다.
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 마찬가지로 10개의 쓰기 고루틴을 띄웁니다.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 1초간 고루틴들의 작업을 대기합니다.
	time.Sleep(time.Second)

	// 마지막으로 각 연산 횟수를 캡쳐하고 리포팅합니다.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
}
