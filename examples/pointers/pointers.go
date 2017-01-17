// Go는 프로그램내에서 값의 참조와 레코드를 전달할 수 있는 <em><a href="http://en.wikipedia.org/wiki/Pointer_(computer_programming)">pointers</a></em>를 지원합니다.

package main

import "fmt"

// 두 함수 `zeroval`와 `zeroptr`를 가지고 포인터가 값과 다르게 어떻게 동작하는지 살펴봅시다.
//  `zeroval`은 `int` 타입 파라미터를 가지므로 인자는 값으로 전달됩니다.
//  `zeroval`는 함수를 호출할 때의 값을 `ival`에 복사하여 가져옵니다.
func zeroval(ival int) {
	ival = 0
}

// 그에 반해 `zeroptr`는 `int`형 포인터를 받도록 `*int` 타입을 파라미터로 갖고 있습니다.
//  함수안에서 `*iptr`는 메모리 주소에서 해당 주소의 현재값으로 포인터를 _역참조(dereference)_ 합니다.
//  역참조된 포인터에 값을 할당하면 이는 참조된 주소의 값을 바꿉니다.
func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	// `&i`는 `i`의 메모리 주소를 반환합니다.
	// 즉, `i`의 포인터입니다.
	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	// 포인터도 출력할 수 있습니다.
	fmt.Println("pointer:", &i)
}
