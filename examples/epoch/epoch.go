// 프로그램에서의 공통적인 요구사항은 [Unix epoch](http://en.wikipedia.org/wiki/Unix_time) 이후로의 초, 밀리초 또는 나노초의 수를 얻는것입니다.
//  Go에서 이를 구하는 방법이 있습니다.

package main

import "fmt"
import "time"

func main() {

	// `time.Now`를 사용하여 유닉스 epoch 이후로 경과된 시간을 각각 초 또는 나노초 단위로 구하기 위해 `Unix` 또는 `UnixNano`를 사용합니다.
	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()
	fmt.Println(now)

	// 참고로, `UnixMillis`는 없기 때문에, 밀리초 단위로 구하려면 나노초값을 수동으로 나누어야합니다.
	millis := nanos / 1000000
	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// epoch 이후의 초 또는 나노초 정수값을 해당하는 `time`으로 변환할 수도 있습니다.
	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))
}
