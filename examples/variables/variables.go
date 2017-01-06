// Go에서, _변수(variables)_는 명시적으로 선언되며 컴파일러에 의해 사용됩니다.
//  예를 들면 함수 호출에서 타입이 정확한지 검사하는데 사용됩니다.

package main

import "fmt"

func main() {

	// `var`는 하나 또는 여러개의 변수를 선언합니다.
	var a string = "initial"
	fmt.Println(a)

	// 한 번에 여러개의 변수를 선언할 수 있습니다.
	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go는 초기화된 변수의 타입을 추론합니다.
	var d = true
	fmt.Println(d)

	// 초기화 없이 선언된 변수는 제로값을 갖게 됩니다.
	//  예를 들어, `int`의 제로값은 `0`입니다.
	var e int
	fmt.Println(e)

	// `:=` 문법은 변수를 선언하는 동시에 초기화하기 위한 단축 표현식입니다.
	//  예를 들면 이 경우는 `var f string = "short"`를 뜻합니다.
	f := "short"
	fmt.Println(f)
}
