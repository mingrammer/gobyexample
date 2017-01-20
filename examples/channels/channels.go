// _Channels_는 동시에 실행되고 있는 고루틴을 연결해주는 파이프입니다.
//  한 고루틴으로부터 채널에 값을 전달하면 다른 고루틴에서 이 값을 받을 수 있습니다.

package main

import "fmt"

func main() {

	// `make(chan val-type)`을 사용해 새로운 채널을 생성합니다.
	//  채널은 전송하는 값의 타입을 타입으로 갖습니다.
	messages := make(chan string)

	// `channel <-` 구문을 사용하여 값을 채널에 _전달합니다(send)_
	//  다음은 새로운 고루틴에서 위에서 만든 `messages` 채널에 `"ping"`을 보냅니다.
	go func() { messages <- "ping" }()

	// `<-channel` 구문은 채널로부터 값을 _받습니다(receives)_.
	//  다음은 위에서 전달된 `"ping"` 메시지를 받아 출력합니다.
	msg := <-messages
	fmt.Println(msg)
}
