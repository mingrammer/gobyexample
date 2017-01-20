// 채널을 함수의 인자로 사용할 때, 채널이 수신용인지 송신용인지를 정할 수 있습니다.
//  이 특성은 프로그램의 타입 안전성을 높여줍니다.

package main

import "fmt"

// `ping` 함수는 송신용 채널만을 받습니다. 이 채널에 수신용 채널을 넣으면 컴파일시 에러가 발생합니다.
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// `pong` 함수는 수신용 채널인 (`pings`)를 하나 받고 송신용인 (`pongs`)를 두 번째 인자로 받습니다.
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "passed message")
	pong(pings, pongs)
	fmt.Println(<-pongs)
}
