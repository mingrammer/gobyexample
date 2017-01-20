// 채널의 송수신은 기본적으로 동기적입니다.
//  그러나, `select`를 `default`문과 함께 사용하면 _non-blocking(비동기)_ 송수신을 구현할 수 있으며, 비동기 다중 `select`도 구현이 가능합니다.

package main

import "fmt"

func main() {
	messages := make(chan string)
	signals := make(chan bool)

	// 다음은 비동기 수신입니다. `messages`에서 값을 사용할 수 있는 경우, `select`는 `<-messages` `case`에서 그 값을 가져옵니다.
	//  그렇지 않은 경우엔 바로 `default` 케이스를 수행합니다.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// 비동기 송신도 유사하게 동작합니다.
	msg := "hi"
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 다중 비동기 select를 구현하기 위해 위의 `default`문에 다중 `case`를 구현할 수 있습니다.
	//  다음은 `messages`와 `signals` 두 채널로부터의 비동기 수신을 시도합니다.
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}
}
