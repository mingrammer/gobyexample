// 기본적으로 채널은  _버퍼가 없는 상태(unbeffered)_이며
//  이는 송신된 값을 받으려고 준비하고 있는 대응되는 리시브 (`<- chan`)가 있는 경우에만 (`chan <-`)을 보낼 수 있음을 의미합니다.
//  _버퍼링된 채널(Buffered channels)_은 대응되는 리시버가 없는 값을 제한된 갯수만큼 허용합니다.

package main

import "fmt"

func main() {

	// 다음은 최대 2개의 값을 버퍼링 할 수 있는 문자열 채널을 `생성(make)`합니다.
	messages := make(chan string, 2)

	// 채널이 버퍼링 되었기 때문에, 동시에 실행중인 대응되는 리시버 없이 채널에 값을 전달할 수 있습니다.
	messages <- "buffered"
	messages <- "channel"

	// 나중에 우리는 평소와 같은 방법으로 이 두 값을 받을 수 있습니다.
	fmt.Println(<-messages)

	fmt.Println(<-messages)
}
