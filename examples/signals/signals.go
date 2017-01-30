// 가끔 우리는 Go 프로그램이 똑똑하게
//  [Unix signal](http://en.wikipedia.org/wiki/Unix_signal)
//  을 처리해주길 바랍니다.
//  예를 들어, `SIGTERM` signal을 받았을 때 적절하게 서버를
//  종료하는 경우나 커맨드라인 도구에서 `SIGINT`를 받았을 때
//  프로세스를 멈추는 경우가 그런 경우입니다.
//  여기서는 채널을 이용하여 signal을 다루는 방법을 살펴보겠습니다.

package main

import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {

	// Go의 signal 알림은 `os.Signal` 값을 채널에 보내는 방식으로 작동합니다.
	//  이런 알림을 받는 채널을 하나 만들어봅시다.
	//  (프로그램을 종료해도 될 때를 알려주는 채널도 만들어볼 것입니다.)
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	// `signal.Notify`는 우리가 지정한 signal을 받을 수 있는
	//  채널을 받고 등록해줍니다.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 이 고루틴은 signal을 받기 위한 블로킹 고루틴입니다.
	//  Signal을 받으면 받은 signal을 출력하고 프로그램에
	//  종료 가능하다고 알려줄 것입니다.
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 프로그램은 여기서 원하는 signal을 받을 때까지 기다릴 것입니다.
	//  (위의 고루틴이 signal을 받으면 `done`에 값을 보내는 것에서
	//  알 수 있습니다.)
	//  값을 받으면, 종료합니다.
	fmt.Println("awaiting signal")
	<-done
	fmt.Println("exiting")
}
