// 이전 예제에서 우린 [atomic operations](/gobyexample/atomic-counters)을 사용해 간단한 카운터 상태를 관리하는 방법을 보았습니다.
//  좀 더 복잡한 상태에 대해서는 여러개의 고루틴이 데이터에 안전하게 접근할 수 있게 _[뮤텍스(mutex)](http://en.wikipedia.org/wiki/Mutual_exclusion)_를 사용할 수 있습니다.

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	// 우리 예제에서 `상태(state)`는 맵입니다.
	var state = make(map[int]int)

	// `mutex`는 `state`에 대한 접근을 동기화합니다.
	var mutex = &sync.Mutex{}

	// 읽기와 쓰기가 얼마나 이루어지는지를 추적할겁니다.
	var readOps uint64 = 0
	var writeOps uint64 = 0

	// 100개의 고루틴을 띄워 각 고루틴에서 1 밀리초마다 반복적으로 읽기를 실행합니다.
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {

				// 각 읽기마다 접근을 위한 키값을 선택합니다.
				//  `state`에 상호 배제 접근을 보장하기 위해 `mutex`를`Lock()`한 뒤 키값으로 값을 읽고, 뮤텍스를 `Unlock()`한 다음 `readOps` 카운트값을 증가시킵니다.
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)

				// 읽기 사이를 조금 대기합니다.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 읽기에서와 같은 패턴으로 쓰기를 시뮬레이션 하기 위해 10개의 고루틴을 띄웁니다.
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// 1초간  `state`와 `mutex`에서 10개의 고루틴 작업을 돌립니다.
	time.Sleep(time.Second)

	// 최종 연산 횟수를 리포팅합니다.
	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	// `state`의 최종 잠금(lock)으로, 어떻게 끝났는지 보여줍니다.
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
