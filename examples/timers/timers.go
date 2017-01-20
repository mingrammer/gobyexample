// 우리는 종종 Go 코드를 나중에 특정 시점에서 실행하거나 일정 간격으로 반복 실행시키고 싶을 때가 있습니다.
//  Go의 내장 기능인 _timer_와 _ticker_는 이를 쉽게 만들어줍니다.
//  먼저 타이머(timers)를 본 뒤, 그 다음 [티커(tickers)](/gobyexample/tickers)에 대해 살펴봅시다.
package main

import "time"
import "fmt"

func main() {

	// 타이머는 미래의 한 이벤트를 나타냅니다. 여러분이 원하는 대기 시간만큼을 타이머에게 지정해주면, 이는 해당 시각에 알림을 주는 채널을 반환합니다.
	//  다음 타이머는 2초를 대기합니다.
	timer1 := time.NewTimer(time.Second * 2)

	// `<-timer1.C`는 타이머가 만료되었음을 알려주는 값을 보내기 전까지 타이머의 `C` 채널을 블로킹합니다.
	<-timer1.C
	fmt.Println("Timer 1 expired")

	// 그저 대기만 하고자 했을땐, `time.Sleep`을 사용할 수 있었습니다.
	//  타이머가 유용한 이유 중 하나는 타이머가 만료되기 전에 취소 시킬 수 있다는 것입니다.
	//  여기에 이에 대한 예시가 있습니다.
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 expired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}
}
