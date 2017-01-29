// [_커맨드라인 인자 (Command-line arguments)_](http://en.wikipedia.org/wiki/Command-line_interface#Arguments)는 프로그램의 실행을 매개변수화하는 일반적인 방법입니다.
//  예를 들어, `go run hello.go`는 `go` 프로그램에 `run`과 `hello.go`를 인자로 사용합니다.

package main

import "os"
import "fmt"

func main() {

	// `os.Args`는 커맨드라인 인자를 그대로 접근하는 방법을 제공합니다.
	//  참고로 이 슬라이스의 첫번째 값은 프로그램의 경로이며 `os.Args[1:]`가 프로그램의 인자를 가집니다.
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	// 일반적인 인덱싱을 사용해 개별 인수들을 얻을 수 있습니다.
	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
