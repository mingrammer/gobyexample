// Go에는 문자열, 정수, 실수, 불리언등을 포함한 다양한 타입들이 존재합니다.
//  여기에 몇 가지 기본적인 예시들이 있습니다.
package main

import "fmt"

func main() {

	// 문자열은 `+`로 합칠 수 있습니다.
	fmt.Println("go" + "lang")

	// 정수와 실수 예시
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0 =", 7.0/3.0)

	// 불리언값들은 여러분들이 예상하는대로 동작합니다.
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)
}
