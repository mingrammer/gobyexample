// Go는 <a href="http://en.wikipedia.org/wiki/Closure_(computer_science)"><em>closures</em></a>를 만들 수 있는
//  [_anonymous functions_](http://en.wikipedia.org/wiki/Anonymous_function)를 지원합니다.
//  익명 함수(Anonymous functions)는 이름 없이 한줄로 함수를 정의하고 싶을때 유용합니다.

package main

import "fmt"

// `intSeq` 함수는 내부에 익명으로 정의한 또 다른 함수를 반환합니다.
//  반환된 함수는 클로저를 만들기 위해 `i` 변수를 `가둬둡니다(closes over).`
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {

	// `intSeq`를 호출하여, 결괏값(함수)을 `nextInt`에 할당합니다.
	//  이 함숫값은 `nextInt`를 호출할 때마다 업데이트 되는 `i` 값을 캡쳐합니다.
	nextInt := intSeq()

	// `netxtInt`를 몇 번 호출하면서 클로저가 어떻게 동작하는지 봅시다.
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 특정 함수에서 상태값이 유일한지 확인하기 위해, 하나를 새롭게 생성하고 테스트 해봅니다.
	newInts := intSeq()
	fmt.Println(newInts())
}
