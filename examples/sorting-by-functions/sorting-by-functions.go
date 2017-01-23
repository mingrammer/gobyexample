// 가끔 일반적인 순서가 아닌 특정한 순서로 컬렉션을 정렬하고 싶을때가 있습니다.
//  예를 들면, 문자열들을 알파벳 순서가 아닌 길이를 기준으로 정렬하고 싶은 경우가 있습니다.
//  이번 장에 Go에서의 커스텀 정렬 예시가 있습니다.

package main

import "sort"
import "fmt"

// Go에서 커스텀 함수로 정렬을 하려면, 그에 해당하는 타입이 필요합니다.
//  우리는 여기서 내장 타입인 `[]string`의 alias인 `ByLength`라는 타입을 만들었습니다.
type ByLength []string

// 우리가 정의한 타입에 `sort.Interface` - `Len`, `Less` 그리고 `Swap`를 구현하면 `sort` 패키지의 제네릭 `Sort` 함수를 사용할 수 있습니다.
//  `Len` 그리고 `Swap`은 일반적으로 타입에 따라 유사하며 `Less`가 실제 커스텀 정렬 로직을 갖습니다.
//  우리의 경우 문자열 길이가 증가하는 순서로 정렬을 하고자 하므로, `len(s[i])`와 `len(s[j])`를 사용합니다.
func (s ByLength) Len() int {
	return len(s)
}
func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

// 구현한 것들을 가지고 원래의 `fruits` 슬라이스를 `ByLength`로 캐스팅하고 여기에 `sort.Sort`를 사용함으로써 우리는 커스텀 정렬을 구현할 수 있습니다.
func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
