// Go의 `sort` 패키지는 내장 타입과 사용자 정의 타입을 위한 정렬을 구현하고 있습니다.
//  먼저 내장 타입에 대한 정렬을 살펴봅시다.

package main

import "fmt"
import "sort"

func main() {

	// 정렬 메서드는 내장 타입에만 해당합니다. 다음은 문자열에 대한 예시입니다.
	//  참고로 정렬은 제자리 정렬(in-place)이므로, 이는 전달된 슬라이스를 변형시키며 새로운 슬라이스를 반환하진 않습니다.
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings:", strs)

	// `int`형의 정렬 예시입니다.
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints:   ", ints)

	// 또한 `sort`를 사용해 슬라이스가 이미 정렬되어 있는지에 대한 여부도 확인할 수 있습니다.
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted: ", s)
}
