// Go에서 _배열(array)_은 특정한 길이의 요소들로 이루어진 순서열입니다.

package main

import "fmt"

func main() {

	// 정확히 5개의 정수를 가진 `a`라는 배열을 만듭니다.
	//  원소의 타입과 길이는 배열 타입의 두 부분입니다.
	//  배열의 기본값은 제로값이며, `int` 배열의 경우 `0` 배열을 의미합니다.
	var a [5]int
	fmt.Println("emp:", a)

	// `array[index] = value` 문법을 사용해 특정 위치에 값을 설정할 수 있으며,
	//  `array[index]`로 값을 가져올 수 있습니다.
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

	// 내장 함수인 `len`은 배열의 길이를 반환합니다.
	fmt.Println("len:", len(a))

	// 한 줄에서 선언과 동시에 배열을 초기화 할때에는 오른쪽 문법을 사용합니다.
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("dcl:", b)

	// 배열 타입은 1차원이지만, 타입들을 구성하여 다차원 자료구조를 만들 수 있습니다.
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
