// _Functions_는 Go의 핵심 기능입니다.
//  몇 가지 예제들을 살펴보면서 함수를 배워봅시다.

package main

import "fmt"

// 여기에 두 `int`를 받아 그의 합을 `int`타입으로 반환하는 함수가 있습니다.
func plus(a int, b int) int {

	// Go에선 명시적 반환을 해야합니다. 즉, 마지막 식의 값을 자동으로 반환하지 않습니다.
	return a + b
}

// 만약 같은 타입을 갖는 파라미터들이 연속적으로 주어지면,
//  마지막 파라미터만 타입을 선언하고 나머지 파라미터들에 대해서는 타입명을 생략할 수 있습니다.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {

	// 함수는 `name(args)`와 같이 호출할 수 있습니다.
	res := plus(1, 2)
	fmt.Println("1+2 =", res)

	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)
}
