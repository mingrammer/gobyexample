// [환경변수](http://en.wikipedia.org/wiki/Environment_variable)
// 는 [Unix 프로그램에 설정 관련 정보를 전달](http://www.12factor.net/config)
// 하는데 일반적으로 사용되는 방법입니다.
//  환경변수를 어떻게 설정하는지, 어떻게 가져오는지, 그리고
//  설정된 환경변수들을 어떻게 보는지 봅시다.

package main

import "os"
import "strings"
import "fmt"

func main() {

	// 환경변수의 키-값 쌍을 설정하려면 `os.Setenv`를 이용합니다.
	//  어떤 키에 매핑된 값은 `os.Getenv`를 통해 가져올 수 있습니다.
	//  이 함수에 주어진 키가 환경변수로 설정되지 않은 키라면
	//  비어있는 문자열을 반환합니다.
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	// `os.Environ` 함수를 이용하여 환경변수로 설정된 모든
	//  키-값 쌍의 리스트를 볼 수 있습니다.
	//  이 함수는 `KEY=value`와 같은 형태의 문자열로 구성된
	//  슬라이스를 반환합니다. 슬라이스의 각 원소에
	//  `string.Split`과 같은 함수를 이용하여
	//  키와 값을 분리할 수 있습니다. 아래는 모든 키를 출력하는 예제입니다.
	fmt.Println()
	for _, e := range os.Environ() {
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
