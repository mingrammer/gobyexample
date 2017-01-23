// 종종 프로그램이 데이터 컬렉션 위에서 작업을 수행해야하는 경우가 있습니다.
//  가령, 주어진 조건을 만족하는 모든 항목을 선택하거나 모든 항목을 커스텀 함수를 가지고 새 컬렉션으로 매핑하는 경우가 있습니다.

// 몇몇 언어에서는 [generic](http://en.wikipedia.org/wiki/Generic_programming) 자료 구조와 알고리즘을 사용하는게 일반적입니다.
//  Go는 제네릭을 지원하지 않습니다. Go에선 여러분의 프로그램과 데이터 타입에 특별히 필요한 경우 컬렉션 함수를 제공하는게 일반적입니다.

// 여기에 `strings`의 슬라이스를 위한 몇가지 컬렉션 함수 예시가 있습니다. 여러분만의 함수를 만들기 위해 이 예시들을 사용할 수 있습니다.
//  어떤 경우엔 헬퍼 함수를 만들고 호출하는 대신 컬렉션 조작 코드를 직접 작성하는게 가장 명확할 수도 있습니다.

package main

import "strings"
import "fmt"

// 문자열 `t`의 첫 위치를 반환하며 매칭이 없을시에는 -1을 반환합니다.
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 문자열 t가 슬라이스에 존재하면 `true`를 반환합니다.
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 슬라이스의 문자열중 하나가 조건 `f`를 만족하면 `true`를 반환합니다.
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 슬라이스의 문자열 모두가 조건 `f`를 만족하면 `true`를 반환합니다.
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 슬라이스에서 조건 `f`를 만족하는 모든 문자열을 포함하는 새로운 슬라이스를 반환합니다.
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 기존 슬라이스에 있는 각각의 문자열에 함수 `f`를 적용한 결괏값들을 포함하는 새로운 슬라이스를 반환합니다.
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

func main() {

	// 다양한 컬렉션 함수를 사용해봅시다.
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(All(strs, func(v string) bool {
		return strings.HasPrefix(v, "p")
	}))

	fmt.Println(Filter(strs, func(v string) bool {
		return strings.Contains(v, "e")
	}))

	// 위의 에시 모두 익명함수를 사용했지만, 올바른 타입을 가진 명명된 함수를 사용할 수도 있습니다.
	fmt.Println(Map(strs, strings.ToUpper))

}
