// Go에서 상태를 관리하는 가장 기본적인 메커니즘은 채널을 통한 통신입니다.
//  이는 [worker pools](/gobyexample/worker-pools) 예제에서 보았습니다.
//  Go에는 상태를 관리하기 위한 몇 가지 방법이 더 있습니다.
//  이번 예제에선 여러개의 고루틴에서 접근되는 _원자성 카운터(atomic counters)_를 위한 `sync/atomic` 패키지를 사용하여 살펴보겠습니다.

package main

import "fmt"
import "time"
import "sync/atomic"

func main() {

	// (항상 양수인) 카운터를 나타내기 위해 부호 없는 정수를 사용합니다.
	var ops uint64 = 0

	// 동시 업데이트를 시뮬레이션 하기 위해, 1 밀리초마다 카운터를 증가시키는 고루틴을 50개 띄웁니다.
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 카운터를 원자적으로 증가시키기 위해 `ops` 카운터의 메모리 주소를 전달하여 `AddUint64`를 사용합니다.
				atomic.AddUint64(&ops, 1)

				// 증가 사이를 조금 대기합니다.
				time.Sleep(time.Millisecond)
			}
		}()
	}

	// ops가 누적될 수 있도록 1초간 대기합니다.
	time.Sleep(time.Second)

	// 카운터가 다른 고루틴에 의해 증가되는 도중에 카운터를 안전하게 사용하기위해, `LoadUint64`를 사용하여 현재값의 복사본을 `opsFinal`로 가져옵니다.
	//  위에서도 보았듯이 값을 페치하기 위해 `&ops`의 메모리 주소를 전달합니다.
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
