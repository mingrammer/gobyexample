// `os.Exit`를 이용하여 프로그램을 지정된 status로 즉시 종료할 수 있습니다.

package main

import "fmt"
import "os"

func main() {

	// `defer`는 `os.Exit`를 이용할 때에는 _작동하지 않습니다._
	//  그래서 이 `fmt.Println`은 절대로 호출되지 않을 것입니다.
	defer fmt.Println("!")

	// status 3으로 종료해봅시다.
	os.Exit(3)
}

// C와 같은 다른 언어와는 달리, Go는 main에서 반환된 값으로
//  exit status를 나타내지 않습니다. 만약 0이 아닌 다른 status로
//  종료하고 싶다면 `os.Exit`를 이용해야 합니다.
