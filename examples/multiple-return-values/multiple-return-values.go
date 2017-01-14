// Go는 자체적으로 _다중 반환값(multiple return values)_을 지원합니다.
//  이는 Go다운 코드를 작성하는데 있어 종종 사용되는데, 예를 들어 함수에서 결괏값과 에러를 동시에 반환하는 경우가 있습니다.

package main

import "fmt"

// 다음 함수 시그니처에 있는 `(int, int)`는 이 함수가 2개의 `int`를 반환한다는걸 의미합니다.
func vals() (int, int) {
	return 3, 7
}

func main() {

	// 다음은 함수 호출로부터 반환되는 두 반환값을 _다중 할당(multiple assignment)_으로 받습니다.
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	// 반환값들중 일부만 사용하고 싶을땐, 공백 식별자 `_`을 사용합니다.
	_, c := vals()
	fmt.Println(c)
}
