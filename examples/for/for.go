// `for` is Gddo's only looping construct. Here are
// three basic types of `for` loops.
// `for`는 Go의 유일한 루프 구조입니다.
//  여기에 기본적인 세 가지 유형의 `for` 루프가 있습니다.

package main

import "fmt"

func main() {

	// The most basic type, with a single condition.
	// 하나의 조건을 갖는 가장 기본적인 유형
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	// 기본적인 초기화/조건/후연산 `for` 루프
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	// 조건이 없는 `for`는 루프 밖으로 `break` 하거나 함수 내부에서 `return`을 하기 전까지 계속 반복합니다.
	for {
		fmt.Println("loop")
		break
	}

	// You can also `continue` to the next iteration of
	// the loop.
	// 루프의 다음 반복을 위해 `continue`를 사용할 수도 있습니다.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
