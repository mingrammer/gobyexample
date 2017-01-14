// [_Variadic functions_](http://en.wikipedia.org/wiki/Variadic_function)는
//  인자 갯수에 상관없이 호출될 수 있습니다.
//  예를 들면, `fmt.Println`가 일반적인 가변 함수입니다.

package main

import "fmt"

// 다음 함수는 `int`형 인자를 임의의 갯수만큼 받습니다.
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// 가변 함수는 일반적인 방법인 개별적인 인자들로 호출할 수 있습니다.
	sum(1, 2)
	sum(1, 2, 3)

	// 만약 값들이 이미 슬라이스에 들어 있다면, 다음과 같이 `func(slice...)`를 사용하여 가변 함수의 인자에 적용하세요.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
