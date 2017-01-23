// _[속도 제한(Rate limiting)](http://en.wikipedia.org/wiki/Rate_limiting)_은 리소스 이용을 제어하고 서비스의 품질을 유지하기위한 중요한 메커니즘입니다.
//  Go는 고루틴, 채널 그리고 [티커(tickers)](/gobyexample/tickers)로 속도 제한을 지원합니다.

package main

import "time"
import "fmt"

func main() {

	// 우선, 기본적인 속도 제한을 살펴봅시다. 요청 핸들링을 제한하고 싶다고 가정해봅시다.
	//  requests라는 이름으로 요청을 처리하는 채널을 만듭니다.
	requests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// 'limiter' 채널은 매 200 밀리초마다 값을 받습니다. 이게 우리 속도 제한 조절기입니다.
	limiter := time.Tick(time.Millisecond * 200)

	// 각 요청을 처리하기 전에 `limiter` 채널의 수신으로 블로킹 함으로써 매 200 밀리초마다 하나의 요청을 받도록 제한하고 있습니다.
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// 전반적으로는 속도 제한을 유지하면서 요청들을 짧게 버스팅(bursts of requests)하고 싶은 경우가 있습니다.
	//  limiter 채널을 버퍼링함으로써 이를 가능케할 수 있습니다.
	//  `burstyLimiter` 채널로 최대 3개의 이벤트를 버스팅할 수 있습니다.
	burstyLimiter := make(chan time.Time, 3)

	// 허용된 버스팅 수만큼 채널을 채웁니다.
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 매 200 밀리초마다 최대 3개까지 `burstyLimiter`에 새로운 값을 추가합니다.
	go func() {
		for t := range time.Tick(time.Millisecond * 200) {
			burstyLimiter <- t
		}
	}()

	// 5개의 요청이 들어오는 시뮬레이션을 해봅시다. 처음 3개의 요청은 `burstyLimiter`의 버스팅 기능의 이점을 얻습니다.
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}
